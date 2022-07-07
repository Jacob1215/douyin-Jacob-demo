package router

import (
	global2 "douyin-Jacob/cmd/api/publish_api/global"
	handlers2 "douyin-Jacob/cmd/api/publish_api/handlers"
	middlewares "douyin-Jacob/pkg/jwt"
	"douyin-Jacob/pkg/tracer"
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
	Publish := Router.Group("/publish").Use(tracer.Trace(global2.ServerConfig.JaegerInfo.Host, global2.ServerConfig.JaegerInfo.Port, global2.ServerConfig.JaegerInfo.Name))
	zap.S().Info("配置publish相关的url")
	{
		Publish.GET("/list", middlewares.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey), handlers2.GetUserVideoList)
		Publish.POST("/action", middlewares.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey), handlers2.PublishVideo)
	}
}