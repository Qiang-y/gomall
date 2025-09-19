package mq

import (
	"github.com/Qiang-y/go-shop/app/frontend/conf"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/nats-io/nats.go"
	"sync"
)

var (
	ConsumerJS nats.JetStreamContext
	ProviderJS nats.JetStreamContext

	Cnc  *nats.Conn
	Pnc  *nats.Conn
	err  error
	once sync.Once
)

func Init() {
	once.Do(func() {
		initProviderJS()
		initConsumerJS()
	})
}

func Close() {
	if Cnc != nil {
		Cnc.Close()
	}
	if Pnc != nil {
		Pnc.Close()
	}
}

func initConsumerJS() {
	Cnc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		hlog.Errorf("Nats Connect Err : %#v", err.Error())
	}
	ConsumerJS, err = Cnc.JetStream()
	if err != nil {
		hlog.Errorf("Build ConsumerJs Err : %#v", err.Error())
	}
}

func initProviderJS() {
	Pnc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		hlog.Errorf("Nats Connect Err : %#v", err.Error())
	}
	//ProviderJS, err = Pnc.JetStream()
	//if err != nil {
	//	hlog.Errorf("Build ConsumerJs Err : %#v", err.Error())
	//}
	//stream, _ := ProviderJS.StreamInfo("Redis_Del")
	//if stream == nil {
	//	stream, err = ProviderJS.AddStream(&nats.StreamConfig{
	//		Name:      "Redis_Del",
	//		Subjects:  []string{"REVIEWS.redis_del"},
	//		Storage:   nats.FileStorage,
	//		Retention: nats.WorkQueuePolicy,
	//	})
	//	if err != nil {
	//		hlog.Errorf("Redis_Del Provider Stream err : %v", err.Error())
	//	}
	//}
}
