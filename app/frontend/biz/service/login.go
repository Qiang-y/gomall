package service

import (
	"biz-demo/gomall/app/frontend/infra/rpc"
	"biz-demo/gomall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"

	auth "biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"
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
	// todo edit your code

	// 调用登录微服务
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
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
