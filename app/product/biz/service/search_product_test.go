package service

import (
	product "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"testing"
)

func TestSearchProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductService(ctx)
	// init req and assert value

	req := &product.SearchProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
