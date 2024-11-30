package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

// Register 注册自定义中间件服务
func Register(h *server.Hertz) {
	h.Use(GlobalAuth())
}
