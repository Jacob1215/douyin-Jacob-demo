package handlers

import (
	global2 "douyin-Jacob/cmd/api/user_api/global"
	"douyin-Jacob/pkg/middleware/models"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"net/http"
	"strconv"
)

type UserInfo struct {
	UserId int64 `form:"User" json:"user_id" binding:"required,max=11"`
	UserName string `form:"username" json:"user_name" binding:"required,max=32"`
	PassWord string `form:"password" json:"pass_word" binding:"required,min=6,max=32"`
	Token string `form:"token" json:"token" binding:"required"`
}

//获取用户信息
func GetUserInfo(c *gin.Context)  {
	UserInfo := models.CustomClaims{}
	userid,_ :=strconv.ParseInt(c.Query("user_id"),10,64)
	token := c.Query("token")
	zap.S().Info(UserInfo.ID,UserInfo.Token)
	//配置熔断限流。
	e,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}
	//获取用户信息
	rsp,err := global2.UserSrvClient.GetUserById(context.WithValue(context.Background(),"ginContext",c),&proto.DouyinUserRequest{
		UserId: userid,
		Token: token,//这个token应该是要验证的。
	});
	if err != nil{
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				SendResponseToHttp(err,c,nil)
			default:
				SendResponseToHttp(err,c,nil)
			}
			return
		}
	}
	e.Exit()
	c.JSON(http.StatusOK,&proto.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg: "get user info success",
		User: rsp.User,
	})
}
