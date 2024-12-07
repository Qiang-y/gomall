package service

import (
	mqutls "biz-demo/gomall/app/frontend/biz/utils"
	cart "biz-demo/gomall/app/frontend/hertz_gen/frontend/cart"
	common "biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"biz-demo/gomall/app/frontend/infra/rpc"
	"biz-demo/gomall/app/frontend/utils"
	rpccart "biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	userId := utils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(userId),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	redisKey := string(userId) + "_cart_num"
	mqutls.SendDelRedisMessage(h.Context, redisKey)
	//err = redis.RedisClient.Del(h.Context, redisKey).Err()
	//if err != nil {
	//	hlog.CtxErrorf(h.Context, "redis Del err : %v", err.Error())
	//}
	return
}
