package api_init

import (
	global2 "douyin-Jacob/cmd/api/relation_api/global"
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

	relationConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port,
			global2.ServerConfig.RelationInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		//只需要就这样使用。设置全局的tracer。但是问题是找不到之前设置的parentSpan的。拿不到父子关系的。
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn]连接【用户服务失败】")
	}
	relationSrvClient := proto.NewRelationSrvClient(relationConn)

	global2.RelationSrvClient = relationSrvClient
}