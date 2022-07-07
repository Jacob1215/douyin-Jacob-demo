package router

import (
	global2 "douyin-Jacob/cmd/api/user_api/global"
	handlers2 "douyin-Jacob/cmd/api/user_api/handlers"
	"douyin-Jacob/pkg/jwt"
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
	InitUserRouter(ApiGroup)
	return Router
}

func InitUserRouter(Router *gin.RouterGroup)  {//注册用户相关的路由
	UserRouter := Router.Group("/user").Use(tracer.Trace(global2.ServerConfig.JaegerInfo.Host, global2.ServerConfig.JaegerInfo.Port, global2.ServerConfig.JaegerInfo.Name))
	zap.S().Info("配置用户相关的url")
	{
		UserRouter.GET("/",middlewares.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey),
			handlers2.GetUserInfo) //先验证,再给Info
		UserRouter.POST("login", handlers2.Login)
		UserRouter.POST("register", handlers2.Register)
	}
}


