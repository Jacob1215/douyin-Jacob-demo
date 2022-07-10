package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/comment_api/global"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListParam struct {
	Token   string `form:"token" json:"token,omitempty"`    // 用户鉴权token
	VideoId int64  `form:"video_id" json:"video_id,omitempty" binding:"required"` // 视频id
}

func CommentList(c *gin.Context)  {
	comListPara := CommentListParam{}
	if err := c.ShouldBind(&comListPara);err !=nil{
		SendHttpResponse(errno.ErrHttpBind,c)
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
			VideoId: comListPara.VideoId,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	sen.Exit()
	c.JSON(http.StatusOK,resp)
}