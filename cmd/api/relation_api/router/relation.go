package router

import (
	"douyin-Jacob/cmd/relation_api/global"
	"douyin-Jacob/cmd/relation_api/handlers"
	"douyin-Jacob/pkg/middleware"
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
	InitRelationRouter(ApiGroup)
	return Router
}

func InitRelationRouter(Router *gin.RouterGroup)  {//注册用户相关的路由
	RelationRouter := Router.Group("/relation").Use(tracer.Trace(global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port,global.ServerConfig.JaegerInfo.Name))
	zap.S().Info("配置Relation相关的url")
	{
		RelationRouter.POST("/action",middlewares.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),
			handlers.RelationAction)//先验证,再给Info
		RelationRouter.GET("/follow_list",middlewares.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),
			handlers.RelationFollowList)//先验证,再给Info
		RelationRouter.GET("/follower_list",middlewares.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),
			handlers.RelationFollowerList)//先验证,再给Info

	}
}


