package service

import (
	"biz-demo/gomall/app/order/biz/dal/mysql"
	"biz-demo/gomall/app/order/biz/model"
	"biz-demo/gomall/rpc_gen/kitex_gen/cart"
	order "biz-demo/gomall/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}

	// 转化信息
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}

		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Street:  v.Consignee.StreetAddress,
				City:    v.Consignee.City,
				State:   v.Consignee.State,
				Country: v.Consignee.Country,
				ZipCode: v.Consignee.ZipCode,
			},
			Items: items,
		})
	}

	// 返回数据
	resp = &order.ListOrderResp{
		Orders: orders,
	}

	return
}
