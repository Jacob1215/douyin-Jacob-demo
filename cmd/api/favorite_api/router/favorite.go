package router

import (
	global2 "douyin-Jacob/cmd/api/favorite_api/global"
	handlers2 "douyin-Jacob/cmd/api/favorite_api/handlers"
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
	InitFavRouter(ApiGroup)
	return Router
}

func InitFavRouter(Router *gin.RouterGroup)  {//注册用户相关的路由
	FavRouter := Router.Group("/favorite").Use(tracer.Trace(global2.ServerConfig.JaegerInfo.Host, global2.ServerConfig.JaegerInfo.Port, global2.ServerConfig.JaegerInfo.Name))
	zap.S().Info("配置favorite相关的url")
	{
		FavRouter.POST("/action",middlewares.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey),
			handlers2.FavoriteAction) //先验证,再给Info
		FavRouter.GET("/list",middlewares.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey), handlers2.FavoriteList)
	}

}