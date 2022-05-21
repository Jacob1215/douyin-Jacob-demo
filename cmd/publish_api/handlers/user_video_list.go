package handlers

import (
	"context"
	"douyin-Jacob/cmd/publish_api/global"
	"douyin-Jacob/cmd/publish_api/models"
	"douyin-Jacob/proto/publish"

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
	zap.S().Info(userInfo.ID,userInfo.Token)
	rsp,err :=global.PublishSrvClient.UserVideoList(context.Background(),&proto.DouyinPublishListRequest{
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
	c.JSON(http.StatusOK,&proto.DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg: "get user video list success",
		VideoList: rsp.VideoList,
	})
}