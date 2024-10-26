package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

const doMutatorName = "do"

var doMutator = object.NewRuleBasedMutator(doMutatorName, slc.Pure(leftVariadicUnaryOp(func(xs []base.Node, x base.Node) optional.Of[base.Node] {
	return optional.Value(x)
})))
