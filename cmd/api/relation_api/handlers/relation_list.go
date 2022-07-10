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

// 用户信息 输出参数
type UserParam struct {
	UserId int64  `form:"user_id" json:"user_id,omitempty" binding:"required"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

func RelationFollowList(c *gin.Context)  {
	var followPara UserParam
	if err := c.ShouldBind(&followPara);err != nil{
		SendHttpResponse(errno.ErrHttpBind,c)
		return
	}
	zap.S().Info(followPara)
	claims,_ := c.Get("claims")
	curUser := claims.(*models.CustomClaims)
	if curUser.Id != followPara.UserId{
		SendHttpResponse(errno.ErrHttpTokenInvalid,c)
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
	resp,err := global2.RelationSrvClient.DouyinRelationFollow(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinRelationFollowListRequest{
			UserId: followPara.UserId,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	e.Exit()
	c.JSON(http.StatusOK,resp)
}

func RelationFollowerList(c *gin.Context)  {
	var followerPara UserParam
	if err := c.ShouldBind(&followerPara);err != nil{
		SendHttpResponse(errno.ErrHttpBind,c)
		return
	}
	claims,_ := c.Get("claims")
	curUser := claims.(*models.CustomClaims)
	if curUser.Id != followerPara.UserId{
		SendHttpResponse(errno.ErrHttpTokenInvalid,c)
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
	resp,err := global2.RelationSrvClient.DouyinRelationFollower(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinRelationFollowerListRequest{
		UserId: followerPara.UserId,
		Token: followerPara.Token,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	e.Exit()
	c.JSON(http.StatusOK,resp)
}