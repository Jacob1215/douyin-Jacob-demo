package feed_init

import (
	global2 "douyin-Jacob/cmd/srv/feed/global"
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
	v.SetConfigFile(constants.FeedSrvConfigFileName)
	if err := v.ReadInConfig();err != nil{
		panic(err)
	}

	//从nacos中读取配置信息
	if err := v.Unmarshal(&global2.NacosConfig);err != nil{
		panic(err)
	}
	zap.S().Info("配置信息:%v", global2.NacosConfig)
	//配置信息使用
	sc := []constant.ServerConfig{
		{
			IpAddr: global2.NacosConfig.Host,
			Port:   global2.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global2.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              constants.LogDir,
		CacheDir:            constants.CacheDir,
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
		DataId: global2.NacosConfig.DataId,
		Group:  global2.NacosConfig.Group,
	})
	if err != nil{panic(err)}
	//json转struct
	err = json.Unmarshal([]byte(content),&global2.ServerConfig)
	if err != nil{
		zap.S().Fatalf("读取nacos配置失败：%s",err)
	}
	fmt.Println(&global2.ServerConfig)

}
