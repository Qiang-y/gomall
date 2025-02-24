package service

import (
	"biz-demo/gomall/app/product/biz/dal/mysql"
	"biz-demo/gomall/app/product/biz/model"
	product "biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
)

type SearchProductService struct {
	ctx context.Context
} // NewSearchProductService new SearchProductService
func NewSearchProductService(ctx context.Context) *SearchProductService {
	return &SearchProductService{ctx: ctx}
}

// Run create note info
func (s *SearchProductService) Run(req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	if err != nil {
		return nil, err
	}
	var results []*product.Product
	for _, v := range products {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
			Quantity:    v.Quantity,
		})
	}
	return &product.SearchProductResp{Results: results}, nil
}
