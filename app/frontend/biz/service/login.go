package service

import (
	"github.com/Qiang-y/go-shop/app/frontend/infra/rpc"
	"github.com/Qiang-y/go-shop/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"

	auth "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 调用登录微服�?
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "/sign-up", err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	err = session.Save()
	if err != nil {
		hlog.CtxErrorf(h.Context, "Session save failed: %v", err)
		return "", err
	}
	hlog.CtxErrorf(h.Context, "Session save success: %v", session.Get("user_id"))
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
