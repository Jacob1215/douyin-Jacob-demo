package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/comment_api/global"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListParam struct {
	Token   string `json:"token,omitempty"`    // 用户鉴权token
	VideoId int64  `json:"video_id,omitempty"` // 视频id
}

func CommentList(c *gin.Context)  {
	var comListPara CommentListParam
	video_id,err := strconv.Atoi(c.Query("video_id"))
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	comListPara.VideoId = int64(video_id)
	comListPara.Token = c.Query("token")
	if len(comListPara.Token) == 0 || comListPara.VideoId < 0 {
		SendResponseToHttp(err,c,nil)
		return
	}
	//配置熔断限流。
	sen,b  := sentinel.Entry("favorite_action",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}


	resp,err := global2.CommentSrvClient.DouyinCommentList(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinCommentListRequest{
			Token: comListPara.Token,
			VideoId: comListPara.VideoId,
		})
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	sen.Exit()
	c.JSON(http.StatusOK,resp)


}