package service

import (
	"biz-demo/gomall/app/product/biz/dal/mysql"
	"biz-demo/gomall/app/product/biz/dal/redis"
	"biz-demo/gomall/app/product/biz/model"
	product "biz-demo/gomall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required") //2004001 表示参数错误
	}
	//productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	productQuery := model.NewCachedProductQuery(s.ctx, mysql.DB, redis.RedisClient)
	p, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Quantity:    p.Quantity,
		},
	}, nil
}
