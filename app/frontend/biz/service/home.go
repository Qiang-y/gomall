package service

import (
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"

	home "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	products, err := rpc.ProductClient.ListProduct(h.Context, &product.ListProductReq{})
	if err != nil {
		return nil, err
	}
	hlog.Infof("ListProduct return ; %v", products.Products)
	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
