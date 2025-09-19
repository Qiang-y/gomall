package cart

import (
	"context"

	"github.com/Qiang-y/go-shop/app/frontend/biz/service"
	"github.com/Qiang-y/go-shop/app/frontend/biz/utils"
	cart "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/cart"
	common "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 放到中间�?
	//// 判断是否登录
	//userId := frontendUtils.GetUserIdFromCtx(ctx)
	//if userId == 0 {
	//	c.Redirect(consts.StatusFound, []byte("/sign-in"))
	//	return
	//}

	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
}

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// 因为是post方法，完成后重定�?
	c.Redirect(consts.StatusFound, []byte("/cart"))
}
