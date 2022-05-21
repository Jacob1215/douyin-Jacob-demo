package router

import (
	"douyin-Jacob/cmd/user_api/global"
	"douyin-Jacob/cmd/user_api/handlers"
	"douyin-Jacob/pkg/middleware"
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
	UserRouter := Router.Group("/user")
	zap.S().Info("配置用户相关的url")
	{
		UserRouter.GET("/",middlewares.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),
			handlers.GetUserInfo)//先验证,再给Info
		UserRouter.POST("login",handlers.Login)
		UserRouter.POST("register",handlers.Register)
	}
}


