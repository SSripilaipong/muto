package global

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

var equalMutator = NewRuleBasedMutatorFromFunctions("==", slc.Pure(strictBinaryOp(func(x, y base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.NodeEqual(x, y)))
})))

var notEqualMutator = NewRuleBasedMutatorFromFunctions("!=", slc.Pure(strictBinaryOp(func(x, y base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.NodeNotEqual(x, y)))
})))

var greaterThanMutator = mutator.MergeNamedSwitches(numberGreaterThanMutator, stringGreaterThanMutator)
var greaterThanOrEqualMutator = mutator.MergeNamedSwitches(numberGreaterThanOrEqualMutator, stringGreaterThanOrEqualMutator)
var lessThanMutator = mutator.MergeNamedSwitches(numberLessThanMutator, stringLessThanMutator)
var lessThanOrEqualMutator = mutator.MergeNamedSwitches(numberLessThanOrEqualMutator, stringLessThanOrEqualMutator)
