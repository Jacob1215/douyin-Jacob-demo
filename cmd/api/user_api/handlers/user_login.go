package handlers

import (
	global2 "douyin-Jacob/cmd/api/user_api/global"
	"douyin-Jacob/pkg/errno"
	middlewares "douyin-Jacob/pkg/jwt"
	"douyin-Jacob/pkg/jwt/models"
	"douyin-Jacob/proto"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"net/http"
	"time"
)

type PassWordLoginForm struct {
	UserName string `form:"username" json:"user_name" binding:"required"`
	PassWord string `form:"password" json:"pass_word" binding:"required,min=6,max=20""`
}


func Login(c *gin.Context)  {
	registerForm := PassWordLoginForm{} //拿到用户注册信息
	//这里为啥一定要用binding呢？
	//表单验证
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	zap.S().Info(registerForm.UserName)
	//配置熔断限流。
	sen,b  := sentinel.Entry("publish_video",sentinel.WithTrafficType(base.Inbound))
	if b !=nil{
		c.JSON(http.StatusTooManyRequests,gin.H{
			"status_code":429,
			"status_msg":"too many requests,please try again later",
		})
		return
	}
	//查询用户存不存在
	if userRsp,err  := global2.UserSrvClient.GetUserInfoByName(context.WithValue(context.Background(),"ginContext",c),&proto.DouyinUserRequest{
		Name: registerForm.UserName,
	});err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				SendHttpResponse(errno.ErrHttpUserNotFound,c)
			default:
				SendHttpResponse(errno.ErrHttpRPCfail,c)
			}
			return
		}
	} else {
		//查询了用户存存在，现在去验证密码
		if passRsp,passErr := global2.UserSrvClient.UserLoginByName(context.WithValue(context.Background(),"ginContext",c),&proto.DouyinUserLoginRequest{
			Password: registerForm.PassWord,
			EncryptedPassword: userRsp.User.Password,
		}); passErr != nil {
			SendHttpResponse(errno.ErrHttpPasswordIncorrect,c)
		} else  {
			if passRsp.StatusCode == 0 {
				//生成token
				j := middlewares.NewJWT(global2.ServerConfig.JWTInfo.SigningKey)
				claims := models.CustomClaims{
					Id: userRsp.User.Id,
					StandardClaims:jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, //180天过期
						Issuer:    "Jacob",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					SendHttpResponse(errno.ErrHttpTokenInvalid,c)
					return
				}
				c.JSON(http.StatusOK, &proto.DouyinUserLoginResponse{
					StatusCode: 0,
					StatusMsg: "Login success",
					UserId: userRsp.User.Id,
					Token: token,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "密码登录失败",
				})
			}
		}
	}
	sen.Exit()
}
