package builtin

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/mutation/object"
)

var stringMutator = object.NewMutator("string", []func(t base.ObjectLike) optional.Of[base.Node]{
	stringOne,
	terminate,
})

func stringOne(t base.ObjectLike) optional.Of[base.Node] {
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
