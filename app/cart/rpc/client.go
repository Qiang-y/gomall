package rpc

import (
	"biz-demo/gomall/app/cart/conf"
	cartUtils "biz-demo/gomall/app/cart/utils"
	"biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	ProductClient productcatalogservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		iniProductClient()
	})
}

func iniProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	cartUtils.MustHandleError(err)
}
