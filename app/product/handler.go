package main

import (
	"biz-demo/gomall/app/product/biz/service"
	product "biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProduct(ctx context.Context, req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	resp, err = service.NewListProductService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProduct(ctx context.Context, req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {
	resp, err = service.NewSearchProductService(ctx).Run(req)

	return resp, err
}

// ReduceProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ReduceProduct(ctx context.Context, req *product.ReduceProductReq) (resp *product.ReduceProductResp, err error) {
	resp, err = service.NewReduceProductService(ctx).Run(req)

	return resp, err
}
