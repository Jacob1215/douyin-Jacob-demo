package service

import (
	"douyin-Jacob/cmd/user/db/model"
	"douyin-Jacob/cmd/user/global"
	"douyin-Jacob/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
)

//通过 name查询用户。
func (s *UserServer)GetUserInfoByName(ctx context.Context,req *proto.DouyinUserRequest)(*proto.DouyinUserResponse,error){
	var user model.User
	zap.S().Infof("%s",req.Name)
	result := global.DB.Where(&model.User{Name: req.Name}).First(&user)
	if result.RowsAffected == 0{
		return nil,status.Errorf(codes.NotFound,"user not exist")
	}
	if result.Error != nil{
		return nil,result.Error
	}
	userInfoRsp := proto.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg: "get user info success",
		User: &proto.UserInfo{
			Id: user.ID,
			Name: user.Name,
			FollowCount: user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow: user.IsFollow,
			Password: user.Password,//这个是加密过后的密码
		},
	}
	return &userInfoRsp,nil
}


//通过Id查询用户
func (s *UserServer)GetUserById(ctx context.Context,req *proto.DouyinUserRequest)(*proto.DouyinUserResponse,error){
	var user model.User
	result := global.DB.First(&user,req.UserId)
	if result.RowsAffected == 0 {
		return nil,status.Errorf(codes.NotFound,"user not exist")
	}
	if result.Error != nil{
		return nil,result.Error
	}
	userInfoRsp := proto.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg: "get user info success",
		User: &proto.UserInfo{
			Id: user.ID,
			Name: user.Name,
			FollowCount: user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow: user.IsFollow,
			Password: user.Password,
		},
	}
	return &userInfoRsp,nil
}
