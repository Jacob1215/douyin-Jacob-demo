package handlers

import (
	global2 "douyin-Jacob/cmd/api/publish_api/global"
	"douyin-Jacob/pkg/middleware/models"
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

type PublishVideoInfo struct {
	Data string `form:"data" json:"data" binding:"required"`
	Token string `form:"token" json:"token" binding:"required"`
	Title string `form:"title" json:"title" binding:"required,max=32"`
}

func PublishVideo(c *gin.Context)  {
	publishvideo := PublishVideoInfo{}
	claims := models.CustomClaims{}

	if err := c.ShouldBindJSON(&publishvideo);err != nil{
		HandleValidatorError(c,err)
		return
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

	zap.S().Infof("%s",publishvideo.Title)

	_,err = global2.PublishSrvClient.PostVideo(context.WithValue(context.Background(),"ginContext",c), //配置tracing
		&proto.DouyinPublishActionRequest{
		User: &proto.User{
			Id:int64(claims.ID),
		},
		Token: publishvideo.Token,
		Title: publishvideo.Title,
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