package service

import (
	"context"
	"douyin-Jacob/cmd/relation/global"
	"douyin-Jacob/dal/db"
	proto "douyin-Jacob/proto"
)

func (s *Relation)DouyinRelationFollow(ctx context.Context, req *proto.DouyinRelationFollowListRequest)(*proto.DouyinRelationFollowListResponse,error)  {
	followUserList := []*db.Relation{}
	err := global.DB.WithContext(ctx).Where("user_id = ?",req.UserId).Find(&followUserList).Error
	if err != nil{
		return nil, err
	}
	followUser := []*db.User{}
	for _,user := range followUserList{
		var user2 *db.User
		if  err = global.DB.WithContext(ctx).First(&user2,user.ID).Error;err != nil{
			return nil,err
		}
		followUser =append(followUser,user2)
	}
	respUserList := []*proto.User{}

	for _,user :=range followUser{
		respUserList =append(respUserList,&proto.User{
			Id: user.ID,
			Name: user.UserName,
			FollowCount: user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow: user.IsFollow,//这个地方TODO，因为不知道拿到的人是否是关注的对象。
		})
	}
	return &proto.DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg: "get Follow list successed",
		UserList: respUserList,
	},nil
}
