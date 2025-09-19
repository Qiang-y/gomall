package main

import (
	"github.com/Qiang-y/go-shop/demo/demo_proto/biz/service"
	pbapi "github.com/Qiang-y/go-shop/demo/demo_proto/kitex_gen/pbapi"
	"context"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
