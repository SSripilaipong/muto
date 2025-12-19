package portal

import (
	"github.com/SSripilaipong/go-common/rods"

	httpserver "github.com/SSripilaipong/muto/builtin/portal/http/server"
	"github.com/SSripilaipong/muto/core/portal"
)

func NewDefaultPortal() *portal.Portal {
	return portal.New(rods.NewMap(map[string]portal.Port{
		"stdout":      NewStdOut(),
		"stdin":       NewStdIn(),
		"spawner":     NewGoroutineSpawner(),
		"chbroker":    NewLocalChannel(),
		"http-server": httpserver.New(),
	}))
}
