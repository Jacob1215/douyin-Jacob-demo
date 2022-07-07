package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)



func Init(Endpoint,OssAccessKeyId,OssSecret string)  *oss.Client{
	client,err := oss.New(Endpoint, OssAccessKeyId, OssSecret)
	if err != nil{
		zap.S().Errorf("Oss client oss failed:%v",err)
	}
	return client
}
