package auth

import (
	"github.com/Qiang-y/go-shop/app/frontend/biz/service"
	"github.com/Qiang-y/go-shop/app/frontend/biz/utils"
	auth "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/auth"
	common "github.com/Qiang-y/go-shop/app/frontend/hertz_gen/frontend/common"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Login .
// @router / [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	redirect, err := service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		//utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		c.Redirect(consts.StatusSeeOther, []byte(redirect))
		return
	}
	c.Redirect(consts.StatusOK, []byte(redirect))
}

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewRegisterService(ctx, c).Run(&req)

	if err != nil {
		//utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		c.Redirect(consts.StatusBadRequest, []byte("/sign-up"))
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))
}

// Logout .
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewLogoutService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))
}
