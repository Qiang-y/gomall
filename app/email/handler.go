package main

import (
	"github.com/Qiang-y/go-shop/app/email/biz/service"
	email "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/email"
	"context"
)

// EmailServiceImpl implements the last service interface defined in the IDL.
type EmailServiceImpl struct{}

// Send implements the EmailServiceImpl interface.
func (s *EmailServiceImpl) Send(ctx context.Context, req *email.EmailReq) (resp *email.EmailResp, err error) {
	resp, err = service.NewSendService(ctx).Run(req)

	return resp, err
}
