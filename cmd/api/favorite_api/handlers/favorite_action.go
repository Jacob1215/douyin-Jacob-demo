package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/favorite_api/global"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/pkg/jwt/models"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)
// 点赞操作 handlers 输入参数
type FavoriteActionParam struct {
	UserId     int64  ` form:"user_id" json:"user_id,omitempty" binding:"required"`     // 用户id
	Token      string ` form:"token" json:"token,omitempty"`       // 用户鉴权token
	VideoId    int64  ` form:"video_id" json:"video_id,omitempty" binding:"required"`    // 视频id
	ActionType int32  ` form:"action_type" json:"action_type,omitempty" binding:"required"` // 1-点赞，2-取消点赞
}
func FavoriteAction(c *gin.Context)  {
	claims,_ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)

	FavActionPar :=  FavoriteActionParam{}
	if err := c.ShouldBind(&FavActionPar);err != nil{
		HandleValidatorError(c,err)
		return
	}
	zap.S().Info(FavActionPar)
	zap.S().Info(currentUser.Id)
	if FavActionPar.UserId != currentUser.Id{
		SendHttpResponse(errno.ErrHttpTokenInvalid,c)
		return
	}
	if FavActionPar.ActionType != 1 && FavActionPar.ActionType != 2{
		SendHttpResponse(errno.ErrHttpInvalidValue,c)
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

	response,err := global2.FavoriteSrvClient.DouyinFavoriteAction(context.WithValue(context.Background(),"ginContext",c),
		&proto.DouyinFavoriteActionRequest{
		VideoId: FavActionPar.VideoId,
		ActionType: FavActionPar.ActionType,
		UserId: FavActionPar.UserId,
		})
	if err != nil{
		SendHttpResponse(errno.ErrHttpRPCfail,c)
		return
	}
	sen.Exit()
	c.JSON(http.StatusOK,response)
}