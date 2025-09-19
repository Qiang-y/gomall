package service

import (
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	frontendUtils "github.com/Qiang-y/go-shop/app/frontend/utils"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	common "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {

	cartResp, err := rpc.CartClient.GetItem(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		hlog.Errorf("GetItem Error : %v", err.Error())
		return nil, err
	}
	var items = make([]map[string]any, 0)
	var total float64
	for _, item := range cartResp.Item {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			// todo: add err resolve
			continue
		}
		p := productResp.Product
		items = append(items, map[string]any{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(p.Price) * float64(item.Quantity)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
