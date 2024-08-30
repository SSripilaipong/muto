package builder

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
)

type Program struct {
	mutate func(object base.ObjectLike) optional.Of[base.Node]
}

func (p Program) InitialObject() base.ObjectLike {
	return base.NewObject(base.NewNamedClass("main"), nil)
}

func (p Program) MutateUntilTerminated(node base.Node) base.Node {
	for base.IsObjectNode(node) {
		newNode := p.mutate(base.UnsafeNodeToObject(node))
		if newNode.IsEmpty() {
			break
		}
		node = newNode.Value()
	}
	return node
}

func (p Program) MutateOnce(node base.Node) base.Node {
	if !base.IsObjectNode(node) {
		return node
	}
	if newNode := p.mutate(base.UnsafeNodeToObject(node)); !newNode.IsEmpty() {
		return newNode.Value()
	}
	return node
}
