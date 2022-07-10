package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/favorite_api/global"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
// 点赞操作 handlers 输入参数
type FavoriteActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	VideoId    int64  `json:"video_id,omitempty"`    // 视频id
	ActionType int32  `json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}
func FavoriteAction(c *gin.Context)  {
	var FavActionPar FavoriteActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	user_id := c.Query("user_id")

	video,err := strconv.Atoi(video_id)
	if err != nil{
		SendHttpResponse(errno.ErrHttpAtoiFail,c)
		return
	}
	action,err := strconv.Atoi(action_type)
	if err != nil{
		SendHttpResponse(errno.ErrHttpAtoiFail,c)
		return
	}
	userId,err := strconv.Atoi(user_id)
	if err != nil{
		SendHttpResponse(errno.ErrHttpAtoiFail,c)
		return
	}

	FavActionPar.Token = token
	FavActionPar.VideoId = int64(video)
	FavActionPar.ActionType = int32(action)
	FavActionPar.UserId = int64(userId)

	//配置熔断限流。
	sen,b  := sentinel.Entry("favorite_action",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}

	response,err := global2.FavoriteSrvClient.DouyinFavoriteAction(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinFavoriteActionRequest{
		VideoId: FavActionPar.VideoId,
		ActionType: FavActionPar.ActionType,
		UserId: FavActionPar.UserId,
		Token: token,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	sen.Exit()
	c.JSON(http.StatusOK,response)
}