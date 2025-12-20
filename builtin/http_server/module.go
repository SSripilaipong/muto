package httpserver

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/module"
	mutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/portal"
	st "github.com/SSripilaipong/muto/syntaxtree"

	httpserverport "github.com/SSripilaipong/muto/builtin/portal/http/server"
)

const ModuleName = "http-server"

type Module struct {
	module.Base
}

func NewModule() Module {
	builder := mutation.NewRuleBuilder()
	normal := slc.Map(builder.BuildNamedUnit)(st.FilterRuleFromStatement(rawStatements))
	collection := mutator.NewCollectionFromMutators(normal, nil)
	return Module{Base: module.NewBase(collection, builder)}
}

func (m Module) MountPortal(q *portal.Portal) {
	m.Base.MountPortal(q)

	port, exists := q.Port("http-server").Return()
	if !exists {
		return
	}
	if mountable, ok := port.(httpserverport.ModuleMountable); ok {
		mountable.MountModule(m)
	}
}
