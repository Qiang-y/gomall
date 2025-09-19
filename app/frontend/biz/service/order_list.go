package service

import (
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	"github.com/Qiang-y/go-shop/app/frontend/types"
	frontendUtils "github.com/Qiang-y/go-shop/app/frontend/utils"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/order"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"time"

	common "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	// 组装订单数据
	var list []types.Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)

		// 组装每个订单商品数据
		for _, vi := range v.Items {
			total += vi.Cost

			i := vi.Item
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: i.ProductId})
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}
			p := productResp.Product

			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Cost:        vi.Cost,
				Qty:         i.Quantity,
			})
		}

		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}

	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
