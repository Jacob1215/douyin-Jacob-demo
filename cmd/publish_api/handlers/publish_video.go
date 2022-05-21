package handlers

import (
	"douyin-Jacob/cmd/publish_api/global"
	"douyin-Jacob/proto/publish"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"context"
	"net/http"
)

type PublishVideoInfo struct {
	Data string `form:"data" json:"data" binding:"required"`
	Token string `form:"token" json:"token" binding:"required"`
	Title string `form:"title" json:"title" binding:"required,max=32"`
}

func PublishVideo(c *gin.Context)  {
	publishvideo := PublishVideoInfo{}
	if err := c.ShouldBindJSON(&publishvideo);err != nil{
		HandleValidatorError(c,err)
	}
	data,_ := json.Marshal(publishvideo.Data)

	zap.S().Infof("%s",publishvideo.Title)
	_,err := global.PublishSrvClient.PostVideo(context.WithValue(context.Background(),"ginContext",c),&proto.DouyinPublishActionRequest{
		Token: publishvideo.Token,
		Title: publishvideo.Title,
		Data: data,
	})
	if err !=nil{
		zap.S().Errorf("failed to publishVideo:%s",err.Error())
		SendResponseToHttp(err,c,nil)
	}
	c.JSON(http.StatusOK,&proto.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg: "Publish video success",
	})

}