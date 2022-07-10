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

// 用户信息 输出参数
type UserParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

func FavoriteList(c *gin.Context)  {
	var paramVar UserParam
	user_id := c.Query("user_id")
	userId,err := strconv.Atoi(user_id)
	if err != nil{
		SendHttpResponse(errno.ErrHttpAtoiFail,c)
		return
	}
	paramVar.UserId = int64(userId)
	paramVar.Token = c.Query("token")
	//配置熔断限流。
	sen,b  := sentinel.Entry("favorite_list",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}

	resp,err := global2.FavoriteSrvClient.DouyinFavoriteList(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinFavoriteListRequest{
			UserId: paramVar.UserId,
			Token: paramVar.Token,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	sen.Exit()
	c.JSON(http.StatusOK,resp)
}