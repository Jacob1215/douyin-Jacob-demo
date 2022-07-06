package service

import (
	"context"
	global2 "douyin-Jacob/cmd/srv/publish/global"
	"douyin-Jacob/dal/db"
	"douyin-Jacob/proto"
	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
//需要对视频列表进行分页，同时返回的是全部视频。
func (s *PublishServer) UserVideoList(ctx context.Context,req *proto.DouyinPublishListRequest)(*proto.DouyinPublishListResponse,error) {
	videoListResponse := proto.DouyinPublishListResponse{}
	var videoList []db.Video
	//拿parentSpan//链路追踪，每个查询服务都去调用一下。
	parentSpan := opentracing.SpanFromContext(ctx)
	userVideoListSpan := opentracing.GlobalTracer().StartSpan("User_video_list",opentracing.ChildOf(parentSpan.Context()))
	result := global2.DB.Where(&db.BaseModel{ID: req.UserId}).Find(&videoList)
	if result.RowsAffected == 0{
		return nil,status.Errorf(codes.NotFound,"User not exist")
	}
	if result.Error != nil{
		return nil,result.Error
	}
	userVideoListSpan.Finish()

	for _,video := range videoList{
		videoListResponse.VideoList = append(videoListResponse.VideoList,&proto.Video{
			Author: &proto.User{
				Id: video.AuthorID,
				Name: video.Author.UserName,
				FollowCount: video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,

				IsFollow: video.Author.IsFollow,
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