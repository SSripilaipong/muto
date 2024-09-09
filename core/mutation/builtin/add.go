package builtin

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/base/datatype"
	"muto/core/mutation/object"
)

const addSymbol = "+"

var addMutator = object.NewMutator(addSymbol, []func(t base.Object) optional.Of[base.Node]{
	addTwo,
})

func addTwo(t base.Object) optional.Of[base.Node] {
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

	newChildren := append([]base.Node{base.NewNumber(c)}, children[2:]...)
	return optional.Value(base.NamedObjectToNode(base.NewNamedObject(rollingAddSymbol, newChildren)))
}
