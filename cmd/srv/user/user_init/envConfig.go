package user_init

import (
	"douyin-Jacob/cmd/user/global"
	"douyin-Jacob/pkg/constants"

	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"fmt"
)


func InitConfig()  {

	v := viper.New()
	v.SetConfigFile(constants.UserSrvConfigFileName)
	if err := v.ReadInConfig();err != nil{
		panic(err)
	}

	//从nacos中读取配置信息
	if err := v.Unmarshal(&global.NacosConfig);err != nil{
		panic(err)
	}
	zap.S().Info("配置信息:%v",global.NacosConfig)
	//配置信息使用
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId: global.NacosConfig.Namespace,
		TimeoutMs: 5000,
		NotLoadCacheAtStart: true,
		LogDir: constants.LogDir,
		CacheDir: constants.CacheDir,
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxAge: 3,
		},
		LogLevel: constants.LogLevel,
	}
	configClient,err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig: &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil{
		panic(err)
	}
	content,err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group: global.NacosConfig.Group,
	})
	if err != nil{panic(err)}
	//json转struct
	err = json.Unmarshal([]byte(content),&global.ServerConfig)
	if err != nil{
		zap.S().Fatalf("读取nacos配置失败：%s",err)
	}
	fmt.Println(&global.ServerConfig)

}
