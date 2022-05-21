package handlers

import (
	"douyin-Jacob/cmd/user_api/global"
	"douyin-Jacob/cmd/user_api/models"
	middlewares "douyin-Jacob/pkg/middleware"
	"douyin-Jacob/proto/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"net/http"
	"time"
)

type PassWordLoginForm struct {
	UserName string `form:"username" json:"user_name" binding:"required,max=32"`
	PassWord string `form:"password" json:"pass_word" binding:"required,min=6,max=32"`
}


func Login(c *gin.Context)  {
	passwordLoginForm:= PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm);err !=nil{
		HandleValidatorError(c,err)
		return
	}

	//查询用户存不存在
	if userRsp,err  := global.UserSrvClient.GetUserInfoByName(context.Background(),&proto.DouyinUserRequest{
		Name: passwordLoginForm.UserName,
	});err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				SendResponseToHttp(err,c,nil)
			default:
				SendResponseToHttp(err,c,nil)
			}
			return
		}
	} else {
		//查询了用户存不存在，现在去验证密码
		if passRsp,passErr := global.UserSrvClient.UserLoginByName(context.Background(),&proto.DouyinUserLoginRequest{
			Password: passwordLoginForm.PassWord,
			EncryptedPassword: userRsp.User.Password,
		}); passErr != nil {
			SendResponseToHttp(err,c,nil)
		} else  {
			if passRsp.StatusCode == 0 {
				//生成token

				j := middlewares.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
				claims := models.CustomClaims{
					ID: uint(userRsp.User.Id),
					StandardClaims:jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
						Issuer:    "Jacob",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					SendResponseToHttp(err,c,nil)
					return
				}
				c.JSON(http.StatusOK, &proto.DouyinUserLoginResponse{
					StatusCode: 0,
					StatusMsg: "Login success",
					UserId: userRsp.User.Id,
					Token: token,
				})
			} else {
				SendResponseToHttp(err,c, nil)
			}
		}
	}
}
