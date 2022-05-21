package handlers

import (
	"douyin-Jacob/cmd/user_api/global"
	"douyin-Jacob/cmd/user_api/models"
	middlewares "douyin-Jacob/pkg/middleware"
	"douyin-Jacob/proto/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"context"
	"net/http"
	"time"
)

type RegisterForm struct {
	UserName string `form:"username" json:"user_name" binding:"required,max=32"`
	PassWord string `form:"password" json:"pass_word" binding:"required,min=6,max=32"`
}

//注册
func Register(c *gin.Context)  {
	registerForm := RegisterForm{}
	if err := c.ShouldBind(&registerForm);err !=nil{
		HandleValidatorError(c,err)
	}
	zap.S().Infof("%s",registerForm.UserName)
	//TODO 要用验证码吗？
	//生成grpc的client接口并调用
	if len(registerForm.PassWord) < 6{ //这个地方还要改，它不是grpc的错误。
		c.JSON(http.StatusBadRequest,&proto.DouyinUserRegisterResponse{
			StatusCode: 400,
			StatusMsg: "Password length must be greater than 6",
		})
		return
	}
	user,err := global.UserSrvClient.UserRegister(context.Background(),&proto.DouyinUserRegisterRequest{
		Username: registerForm.UserName,
		Password: registerForm.PassWord,
	})
	if err !=nil{
		zap.S().Errorf("failed to register:%s",err.Error())
		SendResponseToHttp(err,c,nil)
		return
	}
	//去拿token和验证token
	j := middlewares.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
	cliams := models.CustomClaims{
		ID: uint(user.UserId),
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix(),	//签名的生效时间
			ExpiresAt: time.Now().Unix()+60*60*24*30, // 30天过期
			Issuer: "Jacob", //哪个机构，这个目前写的我的名字。
		},
	}
	token,err := j.CreateToken(cliams)
	if err !=nil{
		SendResponseToHttp(err,c,nil)
		return
	}
	c.JSON(http.StatusOK,&proto.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg: "User register success",
		UserId: user.UserId,
		Token: token,
	})
}