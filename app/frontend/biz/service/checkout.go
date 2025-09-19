package service

import (
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	frontendUtils "github.com/Qiang-y/go-shop/app/frontend/utils"
	rpccart "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	common "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]any
	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetItem(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var total float32

	for _, v := range carts.Item {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}

		p := productResp.Product
		items = append(items, map[string]any{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
