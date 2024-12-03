package service

import (
	"biz-demo/gomall/app/checkout/infra/rpc"
	"biz-demo/gomall/rpc_gen/kitex_gen/cart"
	checkout "biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cartResult, err := rpc.CartClient.GetItem(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Item == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	// 计算购物车内商品总价
	var total float32

	for _, cartItem := range cartResult.Item {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			// todo: 找不到购物车内对应商品的case，暂时先跳过
			continue
		}

		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)
		total += cost
	}

	// 创建订单, 暂时先模拟
	var orderid string
	u, _ := uuid.NewRandom()
	orderid = u.String()

	// 创建支付请求体
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderid,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	// 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	// 调用支付服务
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderid,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
