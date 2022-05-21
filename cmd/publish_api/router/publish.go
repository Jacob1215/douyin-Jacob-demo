package router

import (
	"douyin-Jacob/cmd/publish_api/global"
	"douyin-Jacob/cmd/publish_api/handlers"
	middlewares "douyin-Jacob/pkg/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"net/http"
)

func InitRouters() (c *gin.Engine) {
	Router :=gin.Default()
	//健康检查。
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"success":true,
		})
	})
	ApiGroup := Router.Group("/douyin")
	InitPublishRouter(ApiGroup)
	return Router
}


func InitPublishRouter(Router *gin.RouterGroup)  {
	Publish := Router.Group("/publish")
	zap.S().Info("配置publish相关的url")
	{
		Publish.GET("/list", middlewares.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),handlers.GetUserVideoList)
		Publish.POST("/action", middlewares.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),handlers.PublishVideo)
	}
}