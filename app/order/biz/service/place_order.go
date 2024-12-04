package service

import (
	"biz-demo/gomall/app/order/biz/dal/mysql"
	"biz-demo/gomall/app/order/biz/model"
	order "biz-demo/gomall/rpc_gen/kitex_gen/order"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(500001, "item is empty")
	}
	// 使用事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.Street
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}
		// 写入数据
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var items []model.OrderItem
		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    v.Item.ProductId,
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
			})
		}
		if err := tx.Create(items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})
	if err != nil {
		klog.Error(err)
	}

	return
}
