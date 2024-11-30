package middleware

import (
	frontendUtils "biz-demo/gomall/app/frontend/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
)

// GlobalAuth 全局的登录验证，获得登录的用户ID, 保存到context中往下传
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId != nil {
			hlog.Infof("Authenticated user_id: %v", userId)
		}
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, userId)
		c.Next(ctx)
	}
}

// Auth 鉴定权限
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
