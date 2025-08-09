package module

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

type Dynamic interface {
	Module
	AppendNormal(m mutator.NamedUnit)
	BuildRule(rule st.Rule) mutator.NamedUnit
}
