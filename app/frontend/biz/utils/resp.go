package utils

import (
	"biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "biz-demo/gomall/app/frontend/utils"
	"biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
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
		cartResp, err := rpc.CartClient.GetItem(ctx, &cart.GetCartReq{
			UserId: uint32(frontendUtils.GetUserIdFromCtx(ctx)),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Item)
		}
	}
	return content
}
