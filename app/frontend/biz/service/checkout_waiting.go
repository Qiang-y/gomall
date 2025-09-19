package service

import (
	"github.com/Qiang-y/go-shop/app/frontend/biz/dal/redis"
	checkout "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/checkout"
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	frontendUtils "github.com/Qiang-y/go-shop/app/frontend/utils"
	rpccheckout "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/payment"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardCvv:             req.Cvv,
		},
	})
	if err != nil {
		return nil, err
	}

	redisKey := string(userId) + "_cart_num"
	err = redis.RedisClient.Del(h.Context, redisKey).Err()
	if err != nil {
		hlog.CtxErrorf(h.Context, "redis Del Err : %v", err.Error())
	}

	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
