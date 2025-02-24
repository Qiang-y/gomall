package service

import (
	"biz-demo/gomall/app/product/biz/dal/mysql"
	"biz-demo/gomall/app/product/biz/model"
	product "biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListProductService struct {
	ctx context.Context
} // NewListProductService new ListProductService
func NewListProductService(ctx context.Context) *ListProductService {
	return &ListProductService{ctx: ctx}
}

// Run create note info
func (s *ListProductService) Run(req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
				Quantity:    v.Quantity,
			})
		}
	}
	klog.Debugf("resp products : %v", resp)
	return resp, nil
}
