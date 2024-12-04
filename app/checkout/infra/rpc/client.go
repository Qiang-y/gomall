package rpc

import (
	"biz-demo/gomall/app/checkout/conf"
	"biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client

	once sync.Once
	err  error
)

func InitClient() {
	once.Do(func() {
		iniCartClient()
		iniProductClient()
		iniPaymentClient()
		iniOrderClient()
	})
}

func iniCartClient() {
	var opts []client.Option

	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	CartClient, err = cartservice.NewClient("cart", opts...)
}

func iniProductClient() {
	var opts []client.Option

	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
}

func iniPaymentClient() {
	var opts []client.Option

	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
}

func iniOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
