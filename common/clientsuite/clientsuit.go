package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonClientSuite) Options() []client.Option {
	// 基本配置
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	// 增加链路追踪tracing中间�?
	opts = append(opts, client.WithSuite(tracing.NewClientSuite()))

	// 配置consul服务发现
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))

	return opts
}
