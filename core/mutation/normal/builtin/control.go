package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

const doMutatorName = "do"

var doMutator = NewRuleBasedMutatorFromFunctions(doMutatorName, slc.Pure(leftVariadicUnaryOp(func(xs []base.Node, x base.Node) optional.Of[base.Node] {
	return optional.Value(x)
})))
