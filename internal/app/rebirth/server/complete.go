package server

import (
	"github.com/kataras/iris/v12"
	"github.com/netless-io/rebirth/internal/app/rebirth/hook"
)

func init() {
	app.Handle("POST", "/complete", func(ctx iris.Context) {
		hook.Config.LastStatus = "complete"

		ctx.JSON(struct{}{})
	})
}