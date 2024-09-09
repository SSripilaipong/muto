package builtin

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/object"
)

var concatMutator = object.NewMutator("++", slc.Pure(func(t base.Object) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	n, m := children[0], children[1]
	if !base.IsStringNode(n) || !base.IsStringNode(m) {
		return optional.Empty[base.Node]()
	}
	a, b := base.UnsafeNodeToString(n), base.UnsafeNodeToString(m)
	c := a.Value() + b.Value()

	return valueWithRemainingChildren(base.NewString(c), children[2:])
}))

var stringMutator = object.NewMutator("string", slc.Pure(func(t base.Object) optional.Of[base.Node] {
	children := t.Children()
	if len(children) == 0 {
		return optional.Empty[base.Node]()
	}

	s, isStringer := children[0].(base.MutoStringer)
	if !isStringer {
		return optional.Empty[base.Node]()
	}
	value := base.StringToNode(base.NewString(s.MutoString()))
	return valueWithRemainingChildren(value, children[1:])
}))
