package service

import (
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	rpcproduct "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	p, err := rpc.ProductClient.SearchProduct(h.Context, &rpcproduct.SearchProductReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": p.Results,
		"q":     req.Q,
	}, nil
}
