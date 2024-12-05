package email

import (
	"biz-demo/gomall/app/email/infra/mq"
	"biz-demo/gomall/app/email/infra/notify"
	"biz-demo/gomall/rpc_gen/kitex_gen/email"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	// 为后面接入opentelemetry创建tracer
	tracer := otel.Tracer("shop-nats-consumer-email")

	// 订阅消息
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		// 接入opentelemetry
		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))
		_, span := tracer.Start(ctx, "shop-email-consumer")
		defer span.End()

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
