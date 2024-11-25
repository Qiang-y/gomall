package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:7562"))

	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(http.StatusOK, utils.H{"message": "pong"})
		ctx.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hello world"))
	})

	h.Spin()

}
