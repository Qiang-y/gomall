package consumer

import (
	"biz-demo/gomall/app/frontend/biz/dal/redis"
	"biz-demo/gomall/app/frontend/infra/mq"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/nats-io/nats.go"
)

func Init() {
	mq.Init()
	go initRedisDelConsumer()
}

func initRedisDelConsumer() {
	mq.Cnc.Subscribe("REVIEWS.redis_del", func(msg *nats.Msg) {
		redisKey := string(msg.Data)
		err := redis.RedisClient.Del(context.Background(), redisKey).Err()
		if err != nil {
			hlog.Infof("Redis_Del Key: %s fail", redisKey)
		}

	})
	//_, err := mq.ConsumerJS.Subscribe("REVIEWS.redis_del", func(msg *nats.Msg) {
	//	hlog.Infof("Redis_Del_Consumer recived:%v", msg.Data)
	//
	//	redisKey := string(msg.Data)
	//	err := redis.RedisClient.Del(context.Background(), redisKey).Err()
	//	if err != nil {
	//		hlog.Infof("Redis_Del Key: %s fail", redisKey)
	//		return
	//	}
	//
	//	// 消息确认
	//	if err = msg.Ack(); err != nil {
	//		hlog.Infof("Redis_Del Failed to ack message")
	//	} else {
	//		hlog.Infof("Redis_del successfully : %s", redisKey)
	//	}
	//}, nats.ManualAck(), nats.MaxDeliver(5))
	//if err != nil {
	//	hlog.Error("ConsumerJs Subscribe Err")
	//}

	//// 使用 Close() 关闭所有连接
	//defer func() {
	//	hlog.Info("Closing redis_del nats resources")
	//	mq.Close() // 清理资源
	//}()
}
