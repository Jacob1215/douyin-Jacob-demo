package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/relation_api/global"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/pkg/jwt/models"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 关注操作 handler 输入参数
type RelationActionParam struct {
	UserId     int64  `form:"user_id" json:"user_id,omitempty" binding:"required"`     // 用户id
	Token      string `form:"user_id" json:"token,omitempty" `       // 用户鉴权token
	ToUserId   int64  `form:"user_id" json:"to_user_id,omitempty" binding:"required"`  // 对方用户id
	ActionType int32  `form:"user_id" json:"action_type,omitempty" binding:"required"` // 1-关注，2-取消关注
}

func RelationAction(c *gin.Context)  {
	var relationPara RelationActionParam
	if err := c.ShouldBind(&relationPara);err != nil{
		SendHttpResponse(errno.ErrHttpBind,c)
		return
	}
	claims,_ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	if currentUser.Id != relationPara.UserId {
		SendHttpResponse(errno.ErrHttpTokenInvalid,c)
		return
	}
	if relationPara.ActionType != 1 && relationPara.ActionType != 2{
		SendHttpResponse(errno.ErrHttpInvalidData,c)
		return
	}
	if relationPara.ToUserId == relationPara.UserId {
		SendHttpResponse(errno.ErrHttpInvalidData,c)
		return
	}

	zap.S().Info(relationPara)
	//配置熔断限流。
	e,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}

	resp,err := global2.RelationSrvClient.DouyinRelationAction(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinRelationActionRequest{
			UserId: relationPara.UserId,
			ToUserId: relationPara.ToUserId,
			Token: relationPara.Token,
			ActionType: relationPara.ActionType,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	e.Exit()
	c.JSON(http.StatusOK,resp)
}