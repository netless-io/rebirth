package server

import (
	"github.com/kataras/iris/v12"
	"os"
	"time"
)

type RecordInfo struct {
	UrlAddress string `json:"url"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Filename   string `json:"filename"`
	FPS int `json:"fps"`
}

func RecordInfoAPI(r *RecordInfo, p *os.Process) {
	tid := time.AfterFunc(time.Second*30, func() {
		log.Error("Chrome connection timed out")

		if err := p.Kill(); err != nil {
			log.Fatal("Chrome connection timed out and unable to close chrome. Force quit application")
		}
	})

	app.Handle("POST", "/recordInfo", func(ctx iris.Context) {
		// Turn off the chrome timeout listener
		tid.Stop()

		resp := RecordInfo{
			r.UrlAddress,
			r.StartTime,
			r.EndTime,
			r.Filename,
			r.FPS,
		}

		ctx.JSON(resp)
	})
}
