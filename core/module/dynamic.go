package module

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Dynamic interface {
	Module
	AppendNormal(m mutator.NamedUnit)
	BuildRule(rule st.Rule) mutator.NamedUnit
	BuildNode(node stResult.SimplifiedNode) optional.Of[base.Node]
}
