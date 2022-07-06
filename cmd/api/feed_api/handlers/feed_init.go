package handlers

import (
	"context"
	global2 "douyin-Jacob/cmd/api/feed_api/global"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"

	"github.com/gin-gonic/gin"
	"strconv"

	"net/http"
)

type FeedInfo struct {
	LatestTime string`form:"latestTime" json:"latest_time"`
	Token string `form:"token" json:"token" binding:"required"`
}

func DouyinFeed(c *gin.Context) {

	feedReq := FeedInfo{}
	feedReq.Token = c.Query("token")
	feedReq.LatestTime = c.Query("latest_time")
	var latestTime int64
	if len(feedReq.LatestTime) != 0{
		if latest,err := strconv.Atoi(feedReq.LatestTime);err != nil{
			SendResponseToHttp(err,c,nil)
			return
		}else {
			latestTime = int64(latest)
		}
	}
	//配置熔断限流。
	sen,b  := sentinel.Entry("feed_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}

	videos,err := global2.FeedSrvClient.DouyinFeed(context.WithValue(context.Background(),"ginContext",c),
	&proto.DouyinFeedRequest{
		LatestTime: latestTime,
		Token: feedReq.Token,
	})
	if err != nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	sen.Exit()

	c.JSON(http.StatusOK,&proto.DouyinFeedRespones{
		StatusCode: 0,
		StatusMsg: "get feed videos success",
		VideoList: videos.VideoList,
	})

}