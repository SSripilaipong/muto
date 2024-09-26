package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

const rollingConcatSymbol = "++~"

var rollingConcatMutator = object.NewMutator(rollingConcatSymbol, []func(t base.Object) optional.Of[base.Node]{
	rollingConcatTwo,
	rollingConcatTwoTerminate,
	rollingConcatOne,
	terminate,
})

func rollingConcatTwo(t base.Object) optional.Of[base.Node] {
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
	return optional.Value(base.ObjectToNode(base.NewNamedObject(rollingConcatSymbol, newChildren)))
}

func rollingConcatTwoTerminate(t base.Object) optional.Of[base.Node] {
	children := t.Children()
	if len(children) < 2 {
		return optional.Empty[base.Node]()
	}
	return optional.Value[base.Node](t)
}

func rollingConcatOne(t base.Object) optional.Of[base.Node] {
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
