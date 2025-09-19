package cart

import (
	cart "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart"
	"context"

	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error)
	GetItem(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := cartservice.NewClient(dstService, opts...)
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
	kitexClient cartservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() cartservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error) {
	return c.kitexClient.AddItem(ctx, Req, callOptions...)
}

func (c *clientImpl) GetItem(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	return c.kitexClient.GetItem(ctx, Req, callOptions...)
}

func (c *clientImpl) EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error) {
	return c.kitexClient.EmptyCart(ctx, Req, callOptions...)
}
