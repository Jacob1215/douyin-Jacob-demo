package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/publish_api/global"
	"douyin-Jacob/pkg/jwt/models"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"net/http"
	"strconv"
)

type UserVideoList struct {
	UserId int64 `form:"User" json:"user_id" binding:"required,max=11"`
	Token string `form:"token" json:"token" binding:"required"`
}

//通过用户信息获得user的视频列表
func GetUserVideoList (c *gin.Context)  {
	userInfo := models.CustomClaims{}
	userid,_ :=strconv.ParseInt(c.Query("user_id"),10,64)
	token := c.Query("token")
	zap.S().Info(userInfo.Id,userInfo.AuthorityId)
	//配置熔断限流。
	e,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}
	rsp,err := global2.PublishSrvClient.UserVideoList(context.WithValue(context.Background(),"ginContext",c),&proto.DouyinPublishListRequest{
		UserId: userid,
		Token: token,
	})
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
	c.JSON(http.StatusOK,&proto.DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg: "get user video list success",
		VideoList: rsp.VideoList,
	})
}