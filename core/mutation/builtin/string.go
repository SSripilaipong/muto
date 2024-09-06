package builtin

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/object"
)

var stringMutator = object.NewMutator("string", []func(t base.Object) optional.Of[base.Node]{
	stringOne,
	terminate,
})

func stringOne(t base.Object) optional.Of[base.Node] {
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
}
