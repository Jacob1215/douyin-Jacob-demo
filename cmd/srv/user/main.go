package main

import (
	global2 "douyin-Jacob/cmd/srv/user/global"
	service2 "douyin-Jacob/cmd/srv/user/service"
	user_init2 "douyin-Jacob/cmd/srv/user/user_init"
	"douyin-Jacob/pkg/consul"
	"douyin-Jacob/pkg/initialize"
	"douyin-Jacob/pkg/tracer/otgrpc"
	"douyin-Jacob/pkg/utils"
	"douyin-Jacob/proto"
	"github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	// String defines a string flag with specified name, default value, and usage string.
	// The return value is the address of a string variable that stores the value of the flag.
	Ip := flag.String("ip","0.0.0.0","ip地址")
	Port := flag.Int("port", 50051, "端口号")
	//	initialize
	initialize.InitLogger()
	user_init2.InitConfig()
	user_init2.InitDB()
	// Parse parses the command-line flags from os.Args[1:]. Must be called
	// after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	zap.S().Info("ip:",*Ip)
	*Port,_ =utils.GetFreePort()
	global2.ServerConfig.Port = *Port

	zap.S().Info("port:",*Port)



	//初始化jaeger，
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", global2.ServerConfig.JaegerInfo.Host, global2.ServerConfig.JaegerInfo.Port),
		},
		ServiceName: global2.ServerConfig.JaegerInfo.Name,
	}
	tracer,closer,err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err !=nil{
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	//服务连接建立。//用的grpc
	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
	proto.RegisterUserSrvServer(server,&service2.UserServer{})
	lis,err := net.Listen("tcp",fmt.Sprintf("%s:%d",*Ip,*Port))
	if err != nil{
		panic("failed to listen:"+err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server,health.NewServer())

	//服务注册
	register_client := consul.NewRegistryClient(global2.ServerConfig.ConsulInfo.Host, global2.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s",uuid.NewV4())
	check:=&api.AgentServiceCheck{
		GRPC: fmt.Sprintf("%s:%d", global2.ServerConfig.Host, global2.ServerConfig.Port), //这个端口号一定要改，不然容易出错
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	err = register_client.Register(global2.ServerConfig.Host, global2.ServerConfig.Port, global2.ServerConfig.Name, global2.ServerConfig.Tags, serviceId,check)
	if err != nil {
		zap.S().Panic("服务注册失败:",err.Error())
	}
	zap.S().Debugf("启动服务器，端口：%d", global2.ServerConfig.Port)

	//启动服务
	go func() {
		err = server.Serve(lis)//这个会阻塞，所以放在协程里
		if err != nil{
			panic("failed to start grpc:" + err.Error())
		}
	}()
	//优雅退出,接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	_ = closer.Close()
	if err = register_client.DeRegister(serviceId); err != nil{
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销success")
}
