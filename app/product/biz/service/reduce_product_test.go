package service

import (
	product "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"testing"
)

func TestReduceProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewReduceProductService(ctx)
	// init req and assert value

	req := &product.ReduceProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
