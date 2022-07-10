package service

import (
	"context"
	global2 "douyin-Jacob/cmd/srv/publish/global"
	"douyin-Jacob/dal/db"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/proto"
	"github.com/opentracing/opentracing-go"


)



//需要对视频列表进行分页，同时返回的是全部视频。
func (s *PublishServer) UserVideoList(ctx context.Context,req *proto.DouyinPublishListRequest)(*proto.DouyinPublishListResponse,error) {
	var videoList []*db.Video
	//拿parentSpan//链路追踪，每个查询服务都去调用一下。
	parentSpan := opentracing.SpanFromContext(ctx)
	userVideoListSpan := opentracing.GlobalTracer().StartSpan("User_video_list",opentracing.ChildOf(parentSpan.Context()))
	err := global2.DB.WithContext(ctx).Model(&db.Video{}).Where(&db.Video{AuthorID:req.UserId}).Find(&videoList).Error
	if err != nil{
		return nil,errno.ErrVideoNotFound
	}
	userVideoListSpan.Finish()

	vs ,err := Videos(ctx,videoList,&req.UserId)
	if err != nil{
		return nil, errno.ErrPackVideosErr
	}
	return &proto.DouyinPublishListResponse{
		StatusMsg: "get publish list successed",
		StatusCode: 0,
		VideoList: vs,
	},nil
}