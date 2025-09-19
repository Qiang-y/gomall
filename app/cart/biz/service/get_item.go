package service

import (
	"github.com/Qiang-y/go-shop/app/cart/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/cart/biz/model"
	cart "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetItemService struct {
	ctx context.Context
} // NewGetItemService new GetItemService
func NewGetItemService(ctx context.Context) *GetItemService {
	return &GetItemService{ctx: ctx}
}

// Run create note info
func (s *GetItemService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	list, err := model.GetCartById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var items = make([]*cart.CartItem, 0)
	for _, v := range list {
		items = append(items, &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  v.Qty,
		})
	}
	klog.CtxInfof(s.ctx, "get items : %v", items)
	return &cart.GetCartResp{Item: items}, nil
}
