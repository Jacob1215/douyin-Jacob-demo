package service

import (
	"context"
	"crypto/sha512"
	global2 "douyin-Jacob/cmd/srv/user/global"
	"douyin-Jacob/dal/db"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/proto"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

//通过Name注册。
func (s *UserServer)UserRegister(ctx context.Context,req *proto.DouyinUserRegisterRequest) (*proto.DouyinUserRegisterResponse,error)  {
	zap.S().Infof(req.Username)
	var user db.User
	parentSpan := opentracing.SpanFromContext(ctx)
	userRegisterSpan := opentracing.GlobalTracer().StartSpan("user_register",opentracing.ChildOf(parentSpan.Context()))
	result := global2.DB.Where(&db.User{UserName: req.Username}).First(&user)
	if result.RowsAffected == 1{
		return nil,errno.ErrUserAlreadyExist
	}

	user.UserName = req.Username
	user.FollowCount = 0
	user.FollowerCount = 0
	user.IsFollow = false
	//密码加密。盐值加密
	options := &password.Options{16,100,32,sha512.New}
	salt,encodedPwd := password.Encode(req.Password,options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	result = global2.DB.Create(&user)
	if result.Error != nil{
		return nil,errno.ErrCreateModelErr
	}
	userRegisterSpan.Finish()


	userInfo := &proto.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg: "注册成功",
		UserId: user.ID,
	}
	return userInfo,nil
}