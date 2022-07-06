package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/relation_api/global"
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

func RelationFollowList(c *gin.Context)  {
	var followPara UserParam
	user_id,err := strconv.Atoi(c.Query("user_id"))
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	followPara.UserId = int64(user_id)
	followPara.Token = c.Query("token")

	//配置熔断限流。
	e,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}
	resp,err := global2.RelationSrvClient.DouyinRelationFollow(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinRelationFollowListRequest{
			UserId: followPara.UserId,
			Token:followPara.Token,
		})
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	e.Exit()
	c.JSON(http.StatusOK,resp)
}

func RelationFollowerList(c *gin.Context)  {
	var followerPara UserParam
	user_id,err := strconv.Atoi(c.Query("user_id"))
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	followerPara.UserId = int64(user_id)
	followerPara.Token = c.Query("token")
	//配置熔断限流。
	e,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}
	resp,err := global2.RelationSrvClient.DouyinRelationFollower(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinRelationFollowerListRequest{
		UserId: followerPara.UserId,
		Token: followerPara.Token,
		})
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	e.Exit()
	c.JSON(http.StatusOK,resp)
}