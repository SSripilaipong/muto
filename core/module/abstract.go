package module

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/portal"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Module interface {
	GetClass(name string) base.Class
	AppendNormal(m mutator.NamedUnit)
	BuildRule(rule st.Rule) mutator.NamedUnit
	BuildNode(node stResult.SimplifiedNode) optional.Of[base.Node]
	LoadGlobal(builtin Module)
	MountPortal(q *portal.Portal)
	MutatorCollection() mutator.Collection
}
