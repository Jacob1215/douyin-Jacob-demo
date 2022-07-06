package main

import (
	"douyin-Jacob/cmd/comment_api/global"
	"douyin-Jacob/cmd/comment_api/router"
	"douyin-Jacob/pkg/consul"
	"github.com/hashicorp/consul/api"

	"douyin-Jacob/cmd/comment_api/api_init"
	"douyin-Jacob/pkg/utils"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"os"
	"os/signal"
	"syscall"
)

func main()  {
	//全部初始化
	api_init.InitAll()
	//初始化路由
	Router := router.InitRouters()
	//Viper拦截
	viper.AutomaticEnv()
	debug:=viper.GetBool("debug")
	if !debug{
		_,err:=utils.GetFreePort()
		if err==nil{
			global.ServerConfig.Port=8085 //暂时固定道8087端口。//注意health不成功，可能是端口不空闲。
		}
	}
	//服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port)
	serviceUuid := uuid.NewV4()
	serviceId := fmt.Sprintf("%s",serviceUuid)
	check:=&api.AgentServiceCheck{
		HTTP: fmt.Sprintf("http://%s:%d/health", global.ServerConfig.Host, global.ServerConfig.Port), //这个端口号一定要改，不然容易出错
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	err := register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId,check)
	if err != nil {
		zap.S().Panic("服务注册失败:",err.Error())
	}
	zap.S().Debugf("启动服务器，端口：%d",global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d",global.ServerConfig.Port));err!=nil{
		zap.S().Panic("启动失败：",err.Error())
	}
	//优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit

	if err = register_client.DeRegister(serviceId);err != nil {
		zap.S().Info("注销失败：",err.Error())
	}else{
		zap.S().Info("注销成功：")
	}
}

