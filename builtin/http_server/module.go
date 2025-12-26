package httpserver

import (
	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/module"
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
	mod := global.NewBaseModule()
	normal := slc.Map(mod.BuildRule)(st.FilterRuleFromStatement(rawStatements))

	mod.ExtendCollection(mutator.NewCollectionFromMutators(normal, nil))
	return Module{mod}
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
