package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckoutResultService {
	return &CheckoutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutResultService) Run(req *common.Empty) (resp map[string]any, err error) {
	// 只是返回空页�?
	return utils.H{}, nil
}
