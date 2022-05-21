package main

import (
	"douyin-Jacob/cmd/publish/db"
	"douyin-Jacob/cmd/publish/global"
	"douyin-Jacob/cmd/publish/init"
	"douyin-Jacob/cmd/publish/service"
	"douyin-Jacob/pkg/consul"
	"douyin-Jacob/pkg/initialize"
	"douyin-Jacob/pkg/utils"
	"douyin-Jacob/proto/publish"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"

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
	Port := flag.Int("port",0,"端口号")
	//	initialize
	initialize.InitLogger()
	user_init.InitConfig()
	db.InitDB()
	// Parse parses the command-line flags from os.Args[1:]. Must be called
	// after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	zap.S().Info("ip:",*Ip)
	if *Port == 0{
		*Port,_ = utils.GetFreePort()
	}
	zap.S().Info("port:",*Port)
	global.ServerConfig.Port = *Port
	//服务连接建立。//用的grpc
	server := grpc.NewServer()
	proto.RegisterPublishServer(server,&service.PublishServer{})
	lis,err := net.Listen("tcp",fmt.Sprintf("%s:%d",*Ip,*Port))
	if err != nil{
		panic("failed to listen:"+err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server,health.NewServer())

	//服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s",uuid.NewV4())
	check:=&api.AgentServiceCheck{
		GRPC: fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port), //这个端口号一定要改，不然容易出错
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	err = register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId,check)
	if err != nil {
		zap.S().Panic("服务注册失败:",err.Error())
	}
	zap.S().Debugf("启动服务器，端口：%d",global.ServerConfig.Port)


	if err !=nil{
		panic(err)
	}
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
	if err = register_client.DeRegister(serviceId); err != nil{
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销success")
}