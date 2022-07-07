package main

import (
	api_init2 "douyin-Jacob/cmd/api/publish_api/api_init"
	global2 "douyin-Jacob/cmd/api/publish_api/global"
	router2 "douyin-Jacob/cmd/api/publish_api/router"
	"douyin-Jacob/pkg/consul"
	"github.com/hashicorp/consul/api"

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
	api_init2.InitAll()
	//初始化路由
	Router := router2.InitRouters()
	//Viper拦截
	viper.AutomaticEnv()
	debug:=viper.GetBool("debug")
	if !debug{
		_,err:=utils.GetFreePort()
		if err==nil{
			global2.ServerConfig.Port=8084 //暂时固定道8081端口。
		}
	}
	//服务注册
	register_client := consul.NewRegistryClient(global2.ServerConfig.ConsulInfo.Host, global2.ServerConfig.ConsulInfo.Port)
	serviceUuid := uuid.NewV4()
	serviceId := fmt.Sprintf("%s",serviceUuid)
	check:=&api.AgentServiceCheck{
		HTTP: fmt.Sprintf("http://%s:%d/health", global2.ServerConfig.Host, global2.ServerConfig.Port), //这个端口号一定要改，不然容易出错
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	err := register_client.Register(global2.ServerConfig.Host, global2.ServerConfig.Port, global2.ServerConfig.Name, global2.ServerConfig.Tags, serviceId,check)
	if err != nil {
		zap.S().Panic("服务注册失败:",err.Error())
	}
	zap.S().Debugf("启动服务器，端口：%d", global2.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global2.ServerConfig.Port));err!=nil{
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

