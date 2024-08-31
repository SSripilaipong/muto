package builtin

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/object"
)

var concatMutator = object.NewMutator("++", []func(t base.ObjectLike) optional.Of[base.Node]{
	concatTwo,
	concatTwoTerminate,
	concatOne,
	terminate,
})

func concatTwo(t base.ObjectLike) optional.Of[base.Node] {
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

	newChildren := append([]base.Node{base.NewString(c)}, children[2:]...)
	return optional.Value(base.ObjectToNode(base.NewNamedObject("++", newChildren)))
}

func concatTwoTerminate(t base.ObjectLike) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	return optional.Value[base.Node](t.ConfirmTermination())
}

func concatOne(t base.ObjectLike) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 1 {
		return optional.Empty[base.Node]()
	}
	n := children[0]
	if !base.IsStringNode(n) {
		return optional.Empty[base.Node]()
	}
	return optional.Value(n)
}
