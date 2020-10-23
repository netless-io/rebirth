package server

import (
	"github.com/kataras/iris/v12"
	"github.com/netless-io/rebirth/internal/pkg/logs"
	"os"
)

var log = logs.Server

var app = iris.New()

func Listen(p *os.Process) {
	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:       true,
		DisableInterruptHandler: true,
		Charset:                 "UTF-8",
	})

	if err := app.Listen(":9182", config); err != nil {
		_ = p.Kill()
		log.Fatalf("launch http server failed: %s", err)
	}
}
