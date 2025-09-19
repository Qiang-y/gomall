package rpc

import (
	"github.com/Qiang-y/go-shop/app/frontend/conf"
	frontendUtils "github.com/Qiang-y/go-shop/app/frontend/utils"
	"github.com/Qiang-y/go-shop/common/clientsuite"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client

	once sync.Once
	err  error

	ServiceName  = frontendUtils.ServiceName
	RegistryAddr = conf.GetConf().Hertz.RegistryAddr
)

func Init() {
	once.Do(func() {
		iniUserClient()
		iniProductClient()
		iniCartClient()
		iniCheckoutClient()
		iniOrderClient()
	})
}

func iniUserClient() {
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontendUtils.MustHandleError(err)
	//UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniProductClient() {
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontendUtils.MustHandleError(err)
	//ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniCartClient() {
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontendUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func iniOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}
