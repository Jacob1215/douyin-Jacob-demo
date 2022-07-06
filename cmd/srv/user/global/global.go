package global

import (
	"douyin-Jacob/pkg/config"
	"gorm.io/gorm"
)

//定义全局变量
var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)

