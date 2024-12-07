package service

import (
	"biz-demo/gomall/app/frontend/infra/rpc"
	"biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	category "biz-demo/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	p, err := rpc.ProductClient.ListProduct(h.Context, &product.ListProductReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
