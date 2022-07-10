package handlers

import (
	"douyin-Jacob/cmd/api/publish_api/global"
	"douyin-Jacob/pkg/errno"
	"douyin-Jacob/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (Jwt *jwt.JWT)


func InitJwt()  {
	zap.S().Info(global.ServerConfig.JWTInfo.SigningKey)
	Jwt = jwt.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
}


type Response struct {
	HttpCode int `json:"Http_code"`
	StatusCode int `json:"status_code"`
	StatusMsg string `json:"status_msg"`
}

func SendHttpResponse(err errno.HttpErr,c *gin.Context)  {
	c.JSON(err.StatusCode,Response{
		StatusCode: err.Errno.Code,
		StatusMsg: err.Errno.Message,
	})
}
