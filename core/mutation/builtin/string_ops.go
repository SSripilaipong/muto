package builtin

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/object"
)

var concatMutator = object.NewMutator("++", slc.Pure(binaryOp(func(x, y base.Node) optional.Of[base.Node] {
	if !base.IsStringNode(x) || !base.IsStringNode(y) {
		return optional.Empty[base.Node]()
	}
	a, b := base.UnsafeNodeToString(x), base.UnsafeNodeToString(y)
	return optional.Value[base.Node](base.NewString(a.Value() + b.Value()))
})))

var stringMutator = object.NewMutator("string", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	s, isStringer := x.(base.MutoStringer)
	if !isStringer {
		return optional.Empty[base.Node]()
	}
	return optional.Value[base.Node](base.NewString(s.MutoString()))
})))
