package product

import (
	"context"

	"github.com/Qiang-y/go-shop/app/frontend/biz/service"
	"github.com/Qiang-y/go-shop/app/frontend/biz/utils"
	product "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetProduct .
// @router /product [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, resp))
}

// SearchProduct .
// @router /search [GET]
func SearchProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &common.Empty{}
	resp, err := service.NewSearchProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "search", utils.WarpResponse(ctx, c, resp))
}
