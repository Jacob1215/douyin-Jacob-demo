package oss_init

import (
"douyin-Jacob/pkg/initialize"
)

func InitAll()  {
	//初始化日志
	initialize.InitLogger()
	//初始化配置
	InitConfig()
	//初始化sentinel
	InitSentinel()
	//初始化jaeger
	//tracer.Trace()
}
