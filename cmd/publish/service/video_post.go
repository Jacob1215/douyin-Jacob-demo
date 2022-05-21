package service

import (
	"context"
	"douyin-Jacob/cmd/publish/db/model"
	"douyin-Jacob/cmd/user/global"
	"douyin-Jacob/proto/publish"
)

func (s *PublishServer) PostVideo(ctx context.Context,request *proto.DouyinPublishActionRequest)(*proto.DouyinPublishActionResponse,error) {
	var publishVideo model.VideoPublish
	publishVideo.Title = request.Title
	publishVideo.User.ID = request.User.Id
	publishVideo.Data = request.Data
	global.DB.Save(&publishVideo)
	return &proto.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg: "publish video success",
	},nil
}
