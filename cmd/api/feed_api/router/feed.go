package router

import (
	"douyin-Jacob/cmd/feed_api/global"
	"douyin-Jacob/cmd/feed_api/handlers"

	middlewares "douyin-Jacob/pkg/middleware"
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
	InitFeedRouter(ApiGroup)
	return Router
}

func InitFeedRouter(Router *gin.RouterGroup)  {//注册用户相关的路由
	FeedRouter := Router.Group("/feed").Use(tracer.Trace(global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port,global.ServerConfig.JaegerInfo.Name))
	zap.S().Info("配置feed相关的url")
	{
		FeedRouter.GET("/",middlewares.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),
			handlers.DouyinFeed)//先验证,再给Info
	}
}