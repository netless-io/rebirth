package run

import (
	"github.com/netless-io/rebirth/internal/app/rebirth/chrome"
	"github.com/netless-io/rebirth/internal/app/rebirth/server"
)

type runtimeType int
const (
	none runtimeType = iota
	docker
	local
)
func (e runtimeType) String() string {
	switch e {
	case docker:
		return "docker"
	case local:
		return "local"
	}

	return "none"
}


type runFlags struct {
	RuntimeEnv string
	Record server.RecordInfo
	Chrome chrome.Config
}
