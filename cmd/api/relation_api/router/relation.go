package router

import (
	global2 "douyin-Jacob/cmd/api/relation_api/global"
	handlers2 "douyin-Jacob/cmd/api/relation_api/handlers"
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
	InitRelationRouter(ApiGroup)
	return Router
}

func InitRelationRouter(Router *gin.RouterGroup)  {//注册用户相关的路由
	RelationRouter := Router.Group("/relation").Use(tracer.Trace(global2.ServerConfig.JaegerInfo.Host, global2.ServerConfig.JaegerInfo.Port, global2.ServerConfig.JaegerInfo.Name))
	zap.S().Info("配置Relation相关的url")
	{
		RelationRouter.POST("/action",jwt.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey),
			handlers2.RelationAction) //先验证,再给Info
		RelationRouter.GET("/follow_list",jwt.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey),
			handlers2.RelationFollowList) //先验证,再给Info
		RelationRouter.GET("/follower_list",jwt.JWTAuth(global2.ServerConfig.JWTInfo.SigningKey),
			handlers2.RelationFollowerList) //先验证,再给Info

	}
}


