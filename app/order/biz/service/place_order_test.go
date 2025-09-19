package service

import (
	order "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/order"
	"context"
	"testing"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	req := &order.PlaceOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
