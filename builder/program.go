package builder

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
)

type Program struct {
	mutate func(object base.Object) optional.Of[base.Node]
}

func (p Program) InitialObject() base.Object {
	return base.NewObject(base.NewNamedClass("main"), nil)
}

func (p Program) Mutate(node base.Node) optional.Of[base.Node] {
	if base.IsObjectNode(node) {
		return p.mutate(base.UnsafeNodeToObject(node))
	}
	return optional.Value(node)
}
