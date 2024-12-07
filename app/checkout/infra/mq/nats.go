package mq

import (
	"biz-demo/gomall/app/checkout/conf"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(Nc.Close)
}
