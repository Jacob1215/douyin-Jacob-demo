package service

import (
	"douyin-Jacob/cmd/user/db/model"
	"douyin-Jacob/cmd/user/global"
	"douyin-Jacob/proto/user"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
)

//通过 name查询用户。
func (s *UserServer)GetUserInfoByName(ctx context.Context,req *proto.DouyinUserRequest)(*proto.DouyinUserResponse,error){
	var user model.User
	zap.S().Infof("%s",req.Name)
	parentSpan := opentracing.SpanFromContext(ctx)//这里直接把ctx放进去就好了。//回头写一篇这个文章
	getUserInfoByNameSpan := opentracing.GlobalTracer().StartSpan("get_user_info_by_name",opentracing.ChildOf(parentSpan.Context()))
	result := global.DB.Where(&model.User{Name: req.Name}).First(&user)
	if result.RowsAffected == 0{
		return nil,status.Errorf(codes.NotFound,"user not exist")
	}
	if result.Error != nil{
		return nil,result.Error
	}
	getUserInfoByNameSpan.Finish()
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
	parentSpan := opentracing.SpanFromContext(ctx)
	getUserByIdSpan := opentracing.GlobalTracer().StartSpan("get_user_info_by_Id",opentracing.ChildOf(parentSpan.Context()))
	result := global.DB.First(&user,req.UserId)
	if result.RowsAffected == 0 {
		return nil,status.Errorf(codes.NotFound,"user not exist")
	}
	if result.Error != nil{
		return nil,result.Error
	}
	getUserByIdSpan.Finish()
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
