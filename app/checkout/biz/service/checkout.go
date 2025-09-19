package service

import (
	"github.com/Qiang-y/go-shop/app/checkout/infra/mq"
	"github.com/Qiang-y/go-shop/app/checkout/infra/rpc"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart"
	checkout "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/checkout"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/email"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/order"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/payment"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
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

	var (
		// 计算购物车内商品总价
		total float32
		oi    []*order.OrderItem
	)

	var reduceProd []*product.ReduceProducts

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

		if productResp.Product.Quantity < cartItem.Quantity {
			return nil, kerrors.NewGRPCBizStatusError(5004002, fmt.Sprintf("Product %d have no enough quantity", cartItem.ProductId))
		}

		reduceProd = append(reduceProd, &product.ReduceProducts{
			Id:       cartItem.ProductId,
			Quantity: cartItem.Quantity,
		})

		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)
		total += cost

		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	// 创建订单
	var orderid string
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			Street:  req.Address.StreetAddress,
			City:    req.Address.City,
			State:   req.Address.State,
			Country: req.Address.Country,
			ZipCode: req.Address.ZipCode,
		},
		Items: oi,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	if orderResp != nil && orderResp.Order != nil {
		orderid = orderResp.Order.OrderId
	}

	// 创建支付请求�?
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

	// 清空购物�?
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	// 减少商品库存
	reduceResp, err := rpc.ProductClient.ReduceProduct(s.ctx, &product.ReduceProductReq{Products: reduceProd})
	if !reduceResp.GetSucceed() || err != nil {
		klog.Error(err.Error())
		//return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	// 调用支付服务
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(paymentResult)
	klog.Info(orderResp)

	// 使用Nats消息告知邮件服务发送邮�?
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You have just created an order in the Example Shop",
		Content:     "You have just created an order in the Example Shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}        // 构造mats消息
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header)) // 接入opentelemetry链路追踪
	_ = mq.Nc.PublishMsg(msg)                                                        // 向nats队列发布消息

	resp = &checkout.CheckoutResp{
		OrderId:       orderid,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
