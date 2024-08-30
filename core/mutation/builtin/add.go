package builtin

import (
	"fmt"

	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/base/datatype"
	"phi-lang/core/mutation/object"
	"phi-lang/core/mutation/rule/extractor"
)

func addMutator() object.Mutator {
	return object.NewMutator("+", []func(t extractor.ObjectLike) optional.Of[base.Node]{
		addTwo,
		addOne,
	})
}

func addTwo(t extractor.ObjectLike) optional.Of[base.Node] {
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
	return optional.Value(base.ObjectToNode(base.NewObject(base.NewNamedClass("+"), []base.Node{base.NewNumber(c)})))
}

func addOne(t extractor.ObjectLike) optional.Of[base.Node] {
	fmt.Println("addOne")
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
