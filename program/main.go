package program

import (
	"github.com/SSripilaipong/muto/core/base"
	coreMutation "github.com/SSripilaipong/muto/core/mutation"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type Program struct {
	mutation          coreMutation.TopLevelMutation
	afterMutationHook func(node base.Node)
}

func New(mutation coreMutation.TopLevelMutation) Program {
	return Program{mutation: mutation}
}

func (p Program) InitialObject() base.Object {
	return base.NewNamedOneLayerObject("main", nil)
}

func (p Program) MutateUntilTerminated(node base.Node) base.Node {
	for base.IsMutableNode(node) {
		newNode := p.mutation.Mutate(base.UnsafeNodeToMutable(node))
		if newNode.IsEmpty() {
			break
		}
		p.callAfterMutationHook(newNode.Value())
		node = newNode.Value()
	}
	return node
}

func (p Program) MutateOnce(node base.Node) base.Node {
	if !base.IsMutableNode(node) {
		return node
	}
	newNode := p.mutation.Mutate(base.UnsafeNodeToMutable(node))
	if newNode.IsEmpty() {
		return node
	}
	p.callAfterMutationHook(newNode.Value())
	return newNode.Value()
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

func (p Program) AddRule(rule mutator.NamedObjectMutator) {
	p.mutation.AppendNormalRule(rule)
}
