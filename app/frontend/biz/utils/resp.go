package utils

import (
	frontredis "github.com/Qiang-y/go-shop/app/frontend/biz/dal/redis"
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	frontendUtils "github.com/Qiang-y/go-shop/app/frontend/utils"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"time"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// 添加uerId（判断登录状态）
	userId := frontendUtils.GetUserIdFromCtx(ctx)
	content["user_id"] = userId

	// 判断购物车数�?
	if userId > 0 {
		redisKey := string(userId) + "_cart_num"
		redisCartNum, err := frontredis.RedisClient.Get(ctx, redisKey).Result()
		hlog.CtxInfof(ctx, "redisCartNum = %v", redisCartNum)
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				hlog.CtxErrorf(ctx, "redis get error : %#v", err.Error())
			}
			cartResp, err := rpc.CartClient.GetItem(ctx, &cart.GetCartReq{
				UserId: uint32(userId),
			})
			if err == nil && cartResp != nil {
				content["cart_num"] = len(cartResp.Item)
				frontredis.RedisClient.Set(ctx, redisKey, len(cartResp.Item), 30*time.Second)
			}
		} else if redisCartNum != "0" {
			content["cart_num"] = redisCartNum
		}
	}
	return content
}
