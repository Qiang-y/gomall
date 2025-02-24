package product

import (
	product "biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"

	"biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProduct(ctx context.Context, Req *product.ListProductReq, callOptions ...callopt.Option) (r *product.ListProductResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProduct(ctx context.Context, Req *product.SearchProductReq, callOptions ...callopt.Option) (r *product.SearchProductResp, err error)
	ReduceProduct(ctx context.Context, Req *product.ReduceProductReq, callOptions ...callopt.Option) (r *product.ReduceProductResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProduct(ctx context.Context, Req *product.ListProductReq, callOptions ...callopt.Option) (r *product.ListProductResp, err error) {
	return c.kitexClient.ListProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProduct(ctx context.Context, Req *product.SearchProductReq, callOptions ...callopt.Option) (r *product.SearchProductResp, err error) {
	return c.kitexClient.SearchProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) ReduceProduct(ctx context.Context, Req *product.ReduceProductReq, callOptions ...callopt.Option) (r *product.ReduceProductResp, err error) {
	return c.kitexClient.ReduceProduct(ctx, Req, callOptions...)
}
