package handlers

import (
	"context"
	"douyin-Jacob/cmd/relation_api/global"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 关注操作 handler 输入参数
type RelationActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	ToUserId   int64  `json:"to_user_id,omitempty"`  // 对方用户id
	ActionType int32  `json:"action_type,omitempty"` // 1-关注，2-取消关注
}

func RelationAction(c *gin.Context)  {
	var relationPara RelationActionParam
	token := c.Query("token")
	to_user_id :=c.Query("to_user_id")
	action_type := c.Query("action_type")

	toUserId,err := strconv.Atoi(to_user_id)
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	action ,err :=strconv.Atoi(action_type)
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	relationPara.Token = token
	relationPara.ToUserId = int64(toUserId)
	relationPara.ActionType = int32(action)

	//配置熔断限流。
	e,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}

	resp,err := global.RelationSrvClient.DouyinRelationAction(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinRelationActionRequest{
			ToUserId: relationPara.ToUserId,
			Token: relationPara.Token,
			ActionType: relationPara.ActionType,
		})
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	e.Exit()
	c.JSON(http.StatusOK,resp)
}