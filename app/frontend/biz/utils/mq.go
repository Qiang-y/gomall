package utils

import (
	"biz-demo/gomall/app/frontend/infra/mq"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func SendDelRedisMessage(ctx context.Context, redisKey string) {
	err := mq.Pnc.Publish("REVIEWS.redis_del", []byte(redisKey))
	//_, err := mq.ProviderJS.Publish("redis_del", []byte(redisKey))
	if err != nil {
		hlog.CtxErrorf(ctx, "Redis_Del Provider Send %s err : %v", redisKey, err.Error())
	}
}
