package program

import (
	"github.com/SSripilaipong/muto/core/base"
	coreMutation "github.com/SSripilaipong/muto/core/module"
)

type Program struct {
	mainPackage       coreMutation.Module
	afterMutationHook func(node base.Node)
}

func New(pkg coreMutation.Module) Program {
	return Program{mainPackage: pkg}
}

func (p Program) InitialObject() base.Object {
	return base.NewOneLayerObject(p.mainPackage.GetOrCreateClass("main"), nil)
}

func (p Program) MutateUntilTerminated(node base.Node) base.Node {
	for base.IsMutableNode(node) {
		newNode := base.UnsafeNodeToMutable(node).Mutate()
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
	newNode := base.UnsafeNodeToMutable(node).Mutate()
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

func (p Program) MainPackage() coreMutation.Module {
	return p.mainPackage
}
