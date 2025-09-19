package service

import (
	"github.com/Qiang-y/go-shop/app/product/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/product/biz/dal/redis"
	"github.com/Qiang-y/go-shop/app/product/biz/model"
	product "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ReduceProductService struct {
	ctx context.Context
} // NewReduceProductService new ReduceProductService
func NewReduceProductService(ctx context.Context) *ReduceProductService {
	return &ReduceProductService{ctx: ctx}
}

// Run create note info
func (s *ReduceProductService) Run(req *product.ReduceProductReq) (resp *product.ReduceProductResp, err error) {
	// Finish your business logic.
	if req.Products[0].Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required") //2004001 表示参数错误
	}
	var reduceList []model.ReduceProduct
	for _, item := range req.Products {
		reduceList = append(reduceList, model.ReduceProduct{ID: int(item.GetId()), Quantity: item.GetQuantity()})
	}
	productMutation := model.NewProductMutation(s.ctx, mysql.DB, redis.RedisClient)
	sus, err := productMutation.ReduceQuantity(reduceList)
	if !sus {
		return &product.ReduceProductResp{Succeed: false}, err
	}

	return &product.ReduceProductResp{Succeed: true}, nil
}
