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

func (p Program) Mutate(node base.Node) base.Node {
	if base.IsObjectNode(node) {
		if newNode := p.mutate(base.UnsafeNodeToObject(node)); !newNode.IsEmpty() {
			return newNode.Value()
		}
	}
	return node
}
