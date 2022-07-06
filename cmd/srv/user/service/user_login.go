package service

import (
	"context"
	"crypto/sha512"
	"douyin-Jacob/proto"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/opentracing/opentracing-go"
	"strings"
)

//通过Name登录
func (s *UserServer)UserLoginByName(ctx context.Context,req *proto.DouyinUserLoginRequest)(*proto.DouyinUserLoginResponse,error){
	//校验密码
	options := &password.Options{16,100,32,sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword,"$")
	//zap.S().Infof("加密后的密码：%s",passwordInfo)
	parentSpan := opentracing.SpanFromContext(ctx)
	userLoginSpan := opentracing.GlobalTracer().StartSpan("user_login",opentracing.ChildOf(parentSpan.Context()))
	check := password.Verify(req.Password,passwordInfo[2],passwordInfo[3],options)
	if check  == false{
		return &proto.DouyinUserLoginResponse{
			StatusCode: 1,
			StatusMsg: "password error",
		},nil
	}
	userLoginSpan.Finish()
	return &proto.DouyinUserLoginResponse{
		StatusCode: 0,
	},nil
}