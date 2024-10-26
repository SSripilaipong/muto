package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

var equalMutator = object.NewRuleBasedMutator("==", slc.Pure(binaryOp(func(x, y base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.NodeEqual(x, y)))
})))

var notEqualMutator = object.NewRuleBasedMutator("!=", slc.Pure(binaryOp(func(x, y base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.NodeNotEqual(x, y)))
})))

var greaterThanMutator = object.MergeRuleBasedMutators(numberGreaterThanMutator, stringGreaterThanMutator)
var greaterThanOrEqualMutator = object.MergeRuleBasedMutators(numberGreaterThanOrEqualMutator, stringGreaterThanOrEqualMutator)
var lessThanMutator = object.MergeRuleBasedMutators(numberLessThanMutator, stringLessThanMutator)
var lessThanOrEqualMutator = object.MergeRuleBasedMutators(numberLessThanOrEqualMutator, stringLessThanOrEqualMutator)
