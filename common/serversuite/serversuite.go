package serversuite

import (
	"github.com/Qiang-y/go-shop/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonServerSuite) Options() []server.Option {
	// 配置服务基本信息 �?配置Prometheus链路追踪
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithTracer(prometheus.NewServerTracer("",
			"",
			prometheus.WithDisableServer(true),
			prometheus.WithRegistry(mtl.Registry)),
		),
	}

	// 增加链路追踪tracing中间�?
	opts = append(opts, server.WithSuite(tracing.NewServerSuite()))

	// 将服务注册到consul
	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	return opts
}
