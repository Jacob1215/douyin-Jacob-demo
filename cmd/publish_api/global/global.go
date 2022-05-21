package global

import (
	"douyin-Jacob/pkg/config"
	"douyin-Jacob/proto/publish"
	ut "github.com/go-playground/universal-translator"

)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans ut.Translator
	NacosConfig *config.NacosConfig = &config.NacosConfig{}
	PublishSrvClient proto.PublishClient
)