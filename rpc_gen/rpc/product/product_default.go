package product

import (
	product "biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProduct(ctx context.Context, req *product.ListProductReq, callOptions ...callopt.Option) (resp *product.ListProductResp, err error) {
	resp, err = defaultClient.ListProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProduct(ctx context.Context, req *product.SearchProductReq, callOptions ...callopt.Option) (resp *product.SearchProductResp, err error) {
	resp, err = defaultClient.SearchProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ReduceProduct(ctx context.Context, req *product.ReduceProductReq, callOptions ...callopt.Option) (resp *product.ReduceProductResp, err error) {
	resp, err = defaultClient.ReduceProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ReduceProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
