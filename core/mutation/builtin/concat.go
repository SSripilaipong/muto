package builtin

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/mutation/object"
	"phi-lang/core/mutation/rule/extractor"
)

func concatMutator() object.Mutator {
	return object.NewMutator("++", []func(t extractor.ObjectLike) optional.Of[base.Node]{
		concatTwo,
		concatOne,
	})
}

func concatTwo(t extractor.ObjectLike) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	n, m := children[0], children[1]
	if !base.IsStringNode(n) || !base.IsStringNode(m) {
		return optional.Value[base.Node](t)
	}
	a, b := base.UnsafeNodeToString(n), base.UnsafeNodeToString(m)
	c := a.Value() + b.Value()

	newChildren := append([]base.Node{base.NewString(c)}, children[2:]...)
	return optional.Value(base.ObjectToNode(base.NewObject(base.NewNamedClass("++"), newChildren)))
}

func concatOne(t extractor.ObjectLike) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 1 {
		return optional.Empty[base.Node]()
	}
	n := children[0]
	if !base.IsStringNode(n) {
		return optional.Value[base.Node](t)
	}
	return optional.Value(n)
}
