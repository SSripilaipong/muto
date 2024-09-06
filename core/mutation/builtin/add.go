package builtin

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/base/datatype"
	"muto/core/mutation/object"
)

var addMutator = object.NewMutator("+", []func(t base.ObjectLike) optional.Of[base.Node]{
	addTwo,
	addTwoTerminate,
	addOne,
	terminate,
})

func addTwo(t base.ObjectLike) optional.Of[base.Node] {
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
	return optional.Value(base.NamedObjectToNode(base.NewNamedObject("+", newChildren)))
}

func addTwoTerminate(t base.ObjectLike) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	return optional.Value[base.Node](t.ConfirmTermination())
}

func addOne(t base.ObjectLike) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 1 {
		return optional.Empty[base.Node]()
	}
	n := children[0]
	if !base.IsNumberNode(n) {
		return optional.Empty[base.Node]()
	}
	return optional.Value(n)
}
