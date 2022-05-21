package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"douyin-Jacob/cmd/user_api/global"

	"github.com/gin-gonic/gin"
)

func Trace()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans: true,
				LocalAgentHostPort: fmt.Sprintf("%s:%d",global.ServerConfig.JaegerInfo.Host,global.ServerConfig.JaegerInfo.Port),//默认端口
			},
			ServiceName: global.ServerConfig.JaegerInfo.Name,
		}
		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		//标准输出。// StdLogger is implementation of the Logger interface that delegates to default `log` package
		if err != nil {
			panic(err)
		}
		opentracing.SetGlobalTracer(tracer)//把这个tracer设为全局tracer。其他地方只要import它既可以拿到了。
		defer closer.Close()
		//生成一个startSpan
		startSpan := tracer.StartSpan(ctx.Request.URL.Path)//能拿到ctx里面的一个request的URL了。
		defer startSpan.Finish()
		ctx.Set("tracer",tracer)
		ctx.Set("parentSpan",startSpan)
		ctx.Next()
	}
}