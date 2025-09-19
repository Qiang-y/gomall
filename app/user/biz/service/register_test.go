package service

import (
	"github.com/Qiang-y/go-shop/app/user/biz/dal/mysql"
	user "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/user"
	"context"
	"github.com/joho/godotenv"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "demo@damin.com",
		Password:        "asfdssfd",
		ConfirmPassword: "asfdssfd",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
