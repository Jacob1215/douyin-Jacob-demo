package service

import (
	"context"
	global2 "douyin-Jacob/cmd/srv/comment/global"
	"douyin-Jacob/dal/db"
	proto "douyin-Jacob/proto"
)

func (s *Comment)DouyinCommentList(ctx context.Context,req *proto.DouyinCommentListRequest)(*proto.DouyinCommentListResponse,error){
	comments := []*db.Comment{}
	err := global2.DB.WithContext(ctx).Model(&db.Comment{}).Where(&db.Comment{VideoID: req.VideoId}).Find(&comments).Error
	if err != nil{
		return nil,err
	}
	commentResp := []*proto.Comment{}
	for _,comment := range comments{
		commentResp = append(commentResp,&proto.Comment{
			Id: int64(comment.ID),
			User: &proto.User{
				Id: comment.UserID,
				Name: comment.User.UserName,
				FollowCount: comment.User.FollowCount,
				FollowerCount: comment.User.FollowerCount,
				IsFollow: comment.User.IsFollow,//TODO,这个isfollow是跟relation有关的。
			},
			Content: comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}
	return &proto.DouyinCommentListResponse{
		StatusCode: 0,
		StatusMsg: "get video comments successed",
		CommentList: commentResp,
	},nil
}

