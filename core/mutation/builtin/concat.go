package builtin

import (
	"fmt"

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
		return optional.Empty[base.Node]()
	}
	a, b := base.UnsafeNodeToString(n), base.UnsafeNodeToString(m)
	c := a.Value() + b.Value()
	return optional.Value(base.ObjectToNode(base.NewObject(base.NewNamedClass("++"), []base.Node{base.NewString(c)})))
}

func concatOne(t extractor.ObjectLike) optional.Of[base.Node] {
	fmt.Println("addOne")
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
