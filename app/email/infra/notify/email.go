package notify

import (
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct {
}

// 模拟操作
func (e *NoopEmail) Send(req *email.EmailReq) error {
	pretty.Printf("%v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
