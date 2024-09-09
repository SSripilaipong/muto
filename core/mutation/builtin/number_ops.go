package builtin

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/base/datatype"
	"muto/core/mutation/object"
)

var addMutator = object.NewMutator("+", slc.Pure(binaryOp(func(x, y base.Node) optional.Of[base.Node] {
	if !base.IsNumberNode(x) || !base.IsNumberNode(y) {
		return optional.Empty[base.Node]()
	}
	a, b := base.UnsafeNodeToNumber(x), base.UnsafeNodeToNumber(y)
	return optional.Value[base.Node](base.NewNumber(datatype.AddNumber(a.Value(), b.Value())))
})))

var subtractMutator = object.NewMutator("-", slc.Pure(binaryOp(func(x, y base.Node) optional.Of[base.Node] {
	if !base.IsNumberNode(x) || !base.IsNumberNode(y) {
		return optional.Empty[base.Node]()
	}
	a, b := base.UnsafeNodeToNumber(x), base.UnsafeNodeToNumber(y)
	return optional.Value[base.Node](base.NewNumber(datatype.SubtractNumber(a.Value(), b.Value())))
})))
