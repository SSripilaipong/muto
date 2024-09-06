package builder

import (
	"muto/common/optional"
	"muto/core/base"
)

type Program struct {
	mutate            func(object base.Object) optional.Of[base.Node]
	afterMutationHook func(node base.Node)
}

func (p Program) InitialObject() base.Object {
	return base.NewNamedObject("main", nil)
}

func (p Program) MutateUntilTerminated(node base.Node) base.Node {
	for base.IsNamedObjectNode(node) {
		newNode := p.mutate(base.UnsafeNodeToNamedObject(node))
		if newNode.IsEmpty() {
			break
		}
		p.callAfterMutationHook(newNode.Value())
		node = newNode.Value()
	}
	return node
}

func (p Program) callAfterMutationHook(node base.Node) {
	if p.afterMutationHook != nil {
		p.afterMutationHook(node)
	}
}

func (p Program) WithAfterMutationHook(f func(node base.Node)) Program {
	p.afterMutationHook = f
	return p
}
