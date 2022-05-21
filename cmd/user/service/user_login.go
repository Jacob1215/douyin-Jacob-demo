package service

import (
	"context"
	"crypto/sha512"
	"douyin-Jacob/proto/user"
	"github.com/anaskhan96/go-password-encoder"
	"go.uber.org/zap"
	"strings"
)

//通过Name登录
func (s *UserServer)UserLoginByName(ctx context.Context,req *proto.DouyinUserLoginRequest)(*proto.DouyinUserLoginResponse,error){
	//校验密码
	options := &password.Options{16,100,32,sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword,"$")
	zap.S().Infof("加密后的密码：%s",passwordInfo)
	check := password.Verify(req.Password,passwordInfo[2],passwordInfo[3],options)
	if check  == false{
		return &proto.DouyinUserLoginResponse{
			StatusCode: 1,
			StatusMsg: "password error",
		},nil
	}
	return &proto.DouyinUserLoginResponse{
		StatusCode: 0,
	},nil
}