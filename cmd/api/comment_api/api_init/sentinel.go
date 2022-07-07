package api_init

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"go.uber.org/zap"
)

func InitSentinel()  {
	//初始化sentinel
	err := sentinel.InitDefault()
	if err != nil {
		zap.S().Fatalf("初始化sentinel异常：%v", err)
	}
	//配置限流规则
	//这种配置应该从nacos种读取。这里后面去官方的文档里面修改一下就行了。
	_, err = flow.LoadRules([]*flow.Rule{//可以配置多个
		{
			Resource:               "comment",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,//匀速通过
			Threshold:              3,
			StatIntervalInMs:       6000,
		},
	})
	if err !=nil{
		zap.S().Fatalf("加载规则失败：%v",err)
	}


}