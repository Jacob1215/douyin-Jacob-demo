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
	"strconv"
)

//评论操作， handler输入参数
type CommentActionParam struct {
	UserId      int64   `json:"user_id,omitempty"`      // 用户id
	Token       string  `json:"token,omitempty"`        // 用户鉴权token
	VideoId     int64   `json:"video_id,omitempty"`     // 视频id
	ActionType  int32   `json:"action_type,omitempty"`  // 1-发布评论，2-删除评论
	CommentText *string `json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
}

func CommentAction(c *gin.Context)  {
	var CommentPara CommentActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	video,err := strconv.Atoi(video_id)
	if err != nil{
		SendHttpResponse(errno.ErrHttpVideoNotFound,c)
		return
	}
	action,err := strconv.Atoi(action_type)
	if err != nil{
		SendHttpResponse(errno.ErrHttpAtoiFail,c)
		return
	}
	CommentPara.Token = token
	CommentPara.VideoId = int64(video)
	CommentPara.ActionType = int32(action)
	if action == 1{ //新建评论，就要text就可以了
		comment_text := c.Query("comment_text")
		CommentPara.CommentText = &comment_text
	} else {//删除评论，要comment——id可以。
		comment_id := c.Query("comment_id")
		com_id,err := strconv.Atoi(comment_id)
		if err != nil{
			SendHttpResponse(errno.ErrHttpAtoiFail,c)
			return
		}
		cid64 := int64(com_id)
		CommentPara.CommentId = &cid64
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
			CommentId: *CommentPara.CommentId,
			CommentText: *CommentPara.CommentText,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	sen.Exit()
	c.JSON(http.StatusOK,resp)
}