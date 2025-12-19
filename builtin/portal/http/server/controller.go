package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
)

type controller struct {
	base.NoActiveRule
	server *http.Server
	once   sync.Once
}

func newController(server *http.Server) *base.RuleBasedClass {
	return base.NewRuleBasedClass("<http-server-controller>", &controller{
		server: server,
	})
}

func (r *controller) Normal(obj base.Object) optional.Of[base.Node] {
	return base.StrictTagUnaryOp(func(tag base.Tag) optional.Of[base.Node] {
		if tag.Name() != "stop" {
			return optional.Empty[base.Node]()
		}
		r.once.Do(func() {
			_ = r.server.Shutdown(context.Background())
		})
		return optional.Value[base.Node](base.Null())
	})(obj.ParamChain())
}
