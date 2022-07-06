package api_init

import (
	global2 "douyin-Jacob/cmd/api/comment_api/global"
	"douyin-Jacob/pkg/constants"

	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


func InitConfig()  {

	v := viper.New()//取数据比较简单
	//文件的路径如何设置
	v.SetConfigFile(constants.CommentApiConfigFileName)
	if err := v.ReadInConfig();err!=nil{
		panic(err)
	}

	if err := v.Unmarshal(global2.NacosConfig);err!=nil{
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global2.NacosConfig)

	//viper 的 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化：%v",e.Name)
		_=v.ReadInConfig()
		_=v.Unmarshal(global2.NacosConfig)
		zap.S().Infof("配置信息：%v", global2.NacosConfig)
	})
	//从nacos中读取配置信息
	sc :=[]constant.ServerConfig{
		{
			IpAddr: global2.NacosConfig.Host,
			Port:   global2.NacosConfig.Port,
		},
	}
	cc :=constant.ClientConfig{
		NamespaceId:         global2.NacosConfig.Namespace, //nacos拿的。 //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              constants.LogDir,
		CacheDir:            constants.CacheDir,
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxAge: 3,
		},
		LogLevel:            constants.LogLevel,
	}//其他默认
	// Another way of create config client for dynamic configuration (recommend)
	configClient, err := clients.NewConfigClient(//这里跟老师不一样
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err!=nil{
		panic(err)
	}
	//获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global2.NacosConfig.DataId,
		Group:  global2.NacosConfig.Group})
	if err!=nil{
		panic(err)
	}
	//json转struct。
	err = json.Unmarshal([]byte(content),&global2.ServerConfig)
	if err!=nil{
		zap.S().Fatalf("读取nacos配置失败: %s",err)
	}
	fmt.Println(&global2.ServerConfig)
}






