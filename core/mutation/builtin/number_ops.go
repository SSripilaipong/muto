package builtin

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/base/datatype"
	"muto/core/mutation/object"
)

var addMutator = object.NewMutator("+", slc.Pure(func(t base.Object) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	n, m := children[0], children[1]
	if !base.IsNumberNode(n) || !base.IsNumberNode(m) {
		return optional.Empty[base.Node]()
	}
	a, b := base.UnsafeNodeToNumber(n), base.UnsafeNodeToNumber(m)
	c := datatype.AddNumber(a.Value(), b.Value())

	return valueWithRemainingChildren(base.NewNumber(c), children[2:])
}))

var subtractMutator = object.NewMutator("-", slc.Pure(func(t base.Object) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	n, m := children[0], children[1]
	if !base.IsNumberNode(n) || !base.IsNumberNode(m) {
		return optional.Empty[base.Node]()
	}
	a, b := base.UnsafeNodeToNumber(n), base.UnsafeNodeToNumber(m)
	c := datatype.SubtractNumber(a.Value(), b.Value())

	return valueWithRemainingChildren(base.NewNumber(c), children[2:])
}))
