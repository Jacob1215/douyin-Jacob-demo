package service

import (
	"context"
	"douyin-Jacob/pkg/constants"
	"douyin-Jacob/proto"
	"time"

	"douyin-Jacob/cmd/feed/global"
	"douyin-Jacob/dal/db"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



func (s *FeedSrvServer) DouyinFeed(ctx context.Context,request *proto.DouyinFeedRequest) (*proto.DouyinFeedRespones,error) {
	var videoFeed []db.Video
	var latestTime *int64
	if &request.LatestTime == nil || request.LatestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}

	//TODO 这个时间是怎么处理的。
	res := global.DB.Limit(constants.Limit).Order("update_time desc").Find(&videoFeed, "update_time < ?", time.UnixMilli(*latestTime))
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Video feed not exist")
	}
	if res.Error != nil {
		return nil, res.Error
	}

	//去查找视频的User

	videosListResponse := proto.DouyinFeedRespones{}
	for _,video := range videoFeed{
		videosListResponse.VideoList = append(videosListResponse.VideoList,&proto.Video{
			Author: &proto.User{
				Id: video.AuthorID,
			},
			PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavCount,
			CommentCount: video.ComCount,
			IsFavorite: video.IsFavorite,
			Title: video.Title,
		})
	}
	return &videosListResponse,nil
}
