package service

import (
	"context"
	"douyin-Jacob/cmd/publish/db/model"
	"douyin-Jacob/cmd/publish/global"
	"douyin-Jacob/proto/publish"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
//需要对视频列表进行分页，同时返回的是全部视频。
func (s *PublishServer) UserVideoList(ctx context.Context,req *proto.DouyinPublishListRequest)(*proto.DouyinPublishListResponse,error) {
	videoListResponse := proto.DouyinPublishListResponse{}
	var videoList []model.VideoPublish
	result := global.DB.Where(&model.BaseModel{ID: req.UserId}).Find(&videoList)
	if result.RowsAffected == 0{
		return nil,status.Errorf(codes.NotFound,"User not exist")
	}
	if result.Error != nil{
		return nil,result.Error
	}
	for _,video := range videoList{
		videoListResponse.VideoList = append(videoListResponse.VideoList,&proto.Video{
			VideoId: video.Video,
			Author: &proto.UserIn{
				Id: video.User.ID,
				Name: video.User.Name,
				FollowCount: video.User.FollowCount,
				FollowerCount: video.User.FollowerCount,
				IsFollow: video.User.IsFollow,
			},
			PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavCount,
			CommentCount: video.ComCount,
			IsFavorite: video.IsFavorite,
			Title: video.Title,
		})
	}
	return &videoListResponse,nil
}