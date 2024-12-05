package rpc

import (
	"biz-demo/gomall/app/cart/conf"
	cartUtils "biz-demo/gomall/app/cart/utils"
	"biz-demo/gomall/common/clientsuite"
	"biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err          error
)

func Init() {
	once.Do(func() {
		iniProductClient()
	})
}

func iniProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
