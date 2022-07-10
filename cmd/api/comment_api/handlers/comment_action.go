package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/comment_api/global"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/pkg/jwt/models"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//评论操作， handler输入参数
type CommentActionParam struct {
	UserId      int64   `form:"user_id" json:"user_id,omitempty" binding:"required"`      // 用户id
	Token       string  `form:"token" json:"token,omitempty" `        // 用户鉴权token
	VideoId     int64   `form:"video_id" json:"video_id,omitempty" binding:"required"`     // 视频id
	ActionType  int32   `form:"action_type" json:"action_type,omitempty" binding:"required"`  // 1-发布评论，2-删除评论
	CommentText string `form:"comment_text" json:"comment_text,omitempty" binding:"required"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `form:"comment_id" json:"comment_id,omitempty" `   // 要删除的评论id，在action_type=2的时候使用
}

func CommentAction(c *gin.Context)  {
	CommentPara := CommentActionParam{}
	if err := c.ShouldBind(&CommentPara);err !=nil{
		SendHttpResponse(errno.ErrHttpBind,c)
		return
	}
	claims,_ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	if currentUser.Id  != CommentPara.UserId{
		SendHttpResponse(errno.ErrHttpTokenInvalid,c)
		return
	}
	zap.S().Info(CommentPara)
	if CommentPara.ActionType != 1&& CommentPara.ActionType !=2 {
		SendHttpResponse(errno.ErrHttpInvalidData,c)
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

	resp,err := global2.CommentSrvClient.DouyinCommentAction(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinCommentActionRequest{
			UserId: CommentPara.UserId,
			Token: CommentPara.Token,
			VideoId: CommentPara.VideoId,
			ActionType: CommentPara.ActionType,
			CommentId: CommentPara.CommentId,
			CommentText: CommentPara.CommentText,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	sen.Exit()
	c.JSON(http.StatusOK,resp)
}