package handlers

import (
	global2 "douyin-Jacob/cmd/api/publish_api/global"


	"douyin-Jacob/proto"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bytes"
	"context"
	"io"
	"net/http"
)

func PublishVideo(c *gin.Context)  {

	title := c.PostForm("title")
	token := c.PostForm("token")
	claims,err := Jwt.ParseToken(token)
	if err != nil{
		SendResponseToHttp(err,c,nil)
	}

	data,_,err := c.Request.FormFile("data")
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	defer data.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, data); err != nil {
		SendResponseToHttp(err,c,nil)
		return
	}


	//配置熔断限流。
	e,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}

	_,err = global2.PublishSrvClient.PostVideo(context.WithValue(context.Background(),"ginContext",c), //配置tracing
		&proto.DouyinPublishActionRequest{
		User: &proto.User{
			Id:int64(claims.Id),
		},
		Token: token,
		Title: title,
		Data: buf.Bytes(),
	})
	if err !=nil{
		zap.S().Errorf("failed to publishVideo:%s",err.Error())
		SendResponseToHttp(err,c,nil)
	}
	e.Exit()//管的事以上到限流那块儿的逻辑。

	c.JSON(http.StatusOK,&proto.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg: "Publish video success",
	})

}