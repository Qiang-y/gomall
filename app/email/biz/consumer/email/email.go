package email

import (
	"biz-demo/gomall/app/email/infra/mq"
	"biz-demo/gomall/app/email/infra/notify"
	"biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	// 订阅消息
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})

	if err != nil {
		panic(err)
	}

	// 服务退出前清理Nats
	server.RegisterShutdownHook(func() {
		sub.Unsubscribe() // 取消订阅
		mq.Nc.Close()     // 关闭链接
	})
}
