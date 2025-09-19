package service

import (
	"github.com/Qiang-y/go-shop/app/user/biz/dal/mysql"
	user "github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/user"
	"context"
	"github.com/joho/godotenv"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "demo@damin.com",
		Password: "asfdsfd",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
