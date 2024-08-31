package builder

import (
	"muto/common/optional"
	"muto/core/base"
)

type Program struct {
	mutate func(object base.ObjectLike) optional.Of[base.Node]
}

func (p Program) InitialObject() base.ObjectLike {
	return base.NewNamedObject("main", nil)
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
