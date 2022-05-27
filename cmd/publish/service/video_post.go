package service

import (
	"context"
	"douyin-Jacob/cmd/publish/db/model"
	"douyin-Jacob/cmd/user/global"
	"douyin-Jacob/proto/publish"
	"github.com/opentracing/opentracing-go"
)

func (s *PublishServer) PostVideo(ctx context.Context,request *proto.DouyinPublishActionRequest)(*proto.DouyinPublishActionResponse,error) {
	var publishVideo model.VideoPublish
	publishVideo.Title = request.Title
	publishVideo.User.ID = request.User.Id
	publishVideo.Data = request.Data
	parentSpan := opentracing.SpanFromContext(ctx)
	postVideoSpan := opentracing.GlobalTracer().StartSpan("post_video",opentracing.ChildOf(parentSpan.Context()))

	global.DB.Save(&publishVideo)
	postVideoSpan.Finish()
	return &proto.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg: "publish video success",
	},nil

}
