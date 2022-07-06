package api_init

import (
	global2 "douyin-Jacob/cmd/api/publish_api/global"
	"douyin-Jacob/pkg/tracer/otgrpc"
	"douyin-Jacob/proto"
	"github.com/opentracing/opentracing-go"

	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important//差点没写这个
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//初始化连接
func InitSrvConn()  {
	consulInfo := global2.ServerConfig.ConsulInfo
	zap.S().Info(global2.ServerConfig)
	zap.S().Infof("%s", global2.ServerConfig.PublishSrvInfo)
	publishConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port,
			global2.ServerConfig.PublishSrvInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn]连接【用户服务失败】")
	}
	publishSrvClient :=proto.NewPublishSrvClient(publishConn)
	global2.PublishSrvClient = publishSrvClient
}