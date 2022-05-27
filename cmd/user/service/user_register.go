package service

import (
	"crypto/sha512"
	"douyin-Jacob/cmd/user/db/model"
	"douyin-Jacob/cmd/user/global"
	"douyin-Jacob/proto/user"
	"github.com/opentracing/opentracing-go"

	"github.com/anaskhan96/go-password-encoder"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"fmt"
)

//通过Name注册。
func (s *UserServer)UserRegister(ctx context.Context,req *proto.DouyinUserRegisterRequest) (*proto.DouyinUserRegisterResponse,error)  {
	zap.S().Infof(req.Username)
	var user model.User
	parentSpan := opentracing.SpanFromContext(ctx)
	userRegisterSpan := opentracing.GlobalTracer().StartSpan("user_register",opentracing.ChildOf(parentSpan.Context()))
	result := global.DB.Where(&model.User{Name: req.Username}).First(&user)
	if result.RowsAffected == 1{
		return nil,status.Errorf(codes.AlreadyExists,"用户已存在")
	}

	user.Name = req.Username
	user.FollowCount = 0
	user.FollowerCount = 0
	user.IsFollow = false
	//密码加密。盐值加密
	options := &password.Options{16,100,32,sha512.New}
	salt,encodedPwd := password.Encode(req.Password,options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	result = global.DB.Create(&user)

	userRegisterSpan.Finish()

	if result.Error !=nil{
		return nil,status.Errorf(codes.Internal,result.Error.Error())
	}
	userInfo := &proto.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg: "注册成功",
		UserId: user.ID,

	}
	return userInfo,nil

}