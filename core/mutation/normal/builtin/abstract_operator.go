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
		return base.ProcessMutationResultWithChildren(f(children[0]), children[1:])
	}
}

func binaryOp(f func(x, y base.Node) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return func(t base.Object) optional.Of[base.Node] {
		children := t.Children()
		if len(children) < 2 {
			return optional.Empty[base.Node]()
		}
		return base.ProcessMutationResultWithChildren(f(children[0], children[1]), children[2:])
	}
}
