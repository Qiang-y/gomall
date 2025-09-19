package service

import (
	pbapi "github.com/Qiang-y/go-shop/demo/demo_proto/kitex_gen/pbapi"
	"context"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	// Finish your business logic.
	if info, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME"); ok {
		fmt.Println(info, ok)
	}
	if req.Message == "error" {
		return nil, kerrors.NewGRPCBizStatusError(1004001, "client param error")
	}
	return &pbapi.Response{Message: req.Message}, nil
}
