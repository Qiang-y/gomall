package utils

import (
	frontredis "biz-demo/gomall/app/frontend/biz/dal/redis"
	"biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "biz-demo/gomall/app/frontend/utils"
	"biz-demo/gomall/rpc_gen/kitex_gen/cart"
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

	// 判断购物车数量
	// todo: 用redis替代rpc
	if userId > 0 {
		redisKey := string(userId) + "_cart_num"
		redisCartNum, err := frontredis.RedisClient.Get(ctx, redisKey).Result()
		hlog.CtxInfof(ctx, "redisCartNum = %v", redisCartNum)
		if errors.Is(err, redis.Nil) {
			cartResp, err := rpc.CartClient.GetItem(ctx, &cart.GetCartReq{
				UserId: uint32(userId),
			})
			if err == nil && cartResp != nil {
				content["cart_num"] = len(cartResp.Item)
				frontredis.RedisClient.Set(ctx, redisKey, len(cartResp.Item), 30*time.Second)
			}
		} else if err != nil {
			hlog.CtxErrorf(ctx, "redis get error : %#v", err.Error())
		} else if redisCartNum != "0" {
			content["cart_num"] = redisCartNum
		}
	}
	return content
}
