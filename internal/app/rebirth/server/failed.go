package server

import (
	"github.com/kataras/iris/v12"
	"github.com/netless-io/rebirth/internal/app/rebirth/hook"
)

func init() {
	app.Handle("POST", "/failed", func(ctx iris.Context) {
		hook.Config.LastStatus = "failed"

		ctx.JSON(struct{}{})
	})
}