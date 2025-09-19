package service

import (
	email "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/email"
	"context"
	"testing"
)

func TestSend_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendService(ctx)
	// init req and assert value

	req := &email.EmailReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
