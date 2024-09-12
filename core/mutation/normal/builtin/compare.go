package builtin

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/normal/object"
)

var equalMutator = object.NewMutator("==", slc.Pure(binaryOp(func(x, y base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.NodeEqual(x, y)))
})))

var greaterThanMutator = object.MergeMutators(numberGreaterThanMutator, stringGreaterThanMutator)
var greaterThanOrEqualMutator = object.MergeMutators(numberGreaterThanOrEqualMutator, stringGreaterThanOrEqualMutator)
var lessThanMutator = object.MergeMutators(numberLessThanMutator, stringLessThanMutator)
var lessThanOrEqualMutator = object.MergeMutators(numberLessThanOrEqualMutator, stringLessThanOrEqualMutator)
