package rpc

import (
	"github.com/Qiang-y/go-shop/app/checkout/conf"
	"github.com/Qiang-y/go-shop/common/clientsuite"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client

	once sync.Once
	err  error

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
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
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
}

func iniProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
}

func iniPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
}

func iniOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
