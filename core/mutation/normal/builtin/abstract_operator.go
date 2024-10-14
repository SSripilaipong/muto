package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

func unaryOp(f func(x base.Node) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return func(t base.Object) optional.Of[base.Node] {
		children := t.Children()
		if len(children) == 0 {
			return optional.Empty[base.Node]()
		}
		return processResult(f(children[0]), children[1:])
	}
}

func binaryOp(f func(x, y base.Node) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return func(t base.Object) optional.Of[base.Node] {
		children := t.Children()
		if len(children) < 2 {
			return optional.Empty[base.Node]()
		}
		return processResult(f(children[0], children[1]), children[2:])
	}
}

func processResult(r optional.Of[base.Node], otherChildren []base.Node) optional.Of[base.Node] {
	result, ok := r.Return()
	if !ok {
		return optional.Empty[base.Node]()
	}

	switch {
	case base.IsObjectNode(result):
		obj := base.UnsafeNodeToObject(result)
		return optional.Value[base.Node](obj.AppendChildren(otherChildren))
	default:
		if len(otherChildren) == 0 {
			return optional.Value(result)
		}
		return optional.Value[base.Node](base.NewObject(result, otherChildren))
	}
}
