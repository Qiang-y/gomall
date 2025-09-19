package service

import (
	cart "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart"
	"context"
	"testing"
)

func TestGetItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetItemService(ctx)
	// init req and assert value

	req := &cart.GetCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
