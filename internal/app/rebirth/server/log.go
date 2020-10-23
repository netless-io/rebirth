package server

import (
	"github.com/kataras/iris/v12"
	"github.com/netless-io/rebirth/internal/pkg/logs"
)

func init() {
	type body struct {
		Level string `json:"level"`
		Message string `json:"message"`
	}

	app.Handle("POST", "/log", func(ctx iris.Context) {
		var b body

		if err := ctx.ReadJSON(&b); err != nil {
			ctx.JSON(struct{}{})
			return
		}

		log := logs.Extension

		switch b.Level {
		case "debug":
			log.Debugf(b.Message)
		case "info":
			log.Infoln(b.Message)
		case "warn":
			log.Warnln(b.Message)
		case "error":
			log.Errorln(b.Message)
		default:
			log.Infoln(b.Message)
		}

		ctx.JSON(struct{}{})
	})
}