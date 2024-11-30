package service

import (
	"biz-demo/gomall/app/frontend/infra/rpc"
	rpcproduct "biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "biz-demo/gomall/app/frontend/hertz_gen/frontend/product"
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
	// todo edit your code

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
