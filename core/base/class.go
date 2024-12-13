package base

import (
	"github.com/SSripilaipong/muto/common/optional"
)

type Class struct {
	name string
}

func (c Class) Children() []Node {
	return nil
}

func (c Class) NodeType() NodeType {
	return NodeTypeClass
}

func (c Class) Mutate(mutation NameWiseMutation) optional.Of[Node] {
	return c.MutateAsHead(nil, mutation)
}

func (c Class) MutateAsHead(children []Node, mutation NameWiseMutation) optional.Of[Node] {
	if result, ok := c.ActivelyMutateWithObjMutateFunc(children, mutation).Return(); ok {
		return optional.Value(result)
	}
	return c.MutateWithObjMutateFunc(children, mutation)
}

func (c Class) ActivelyMutateWithObjMutateFunc(children []Node, mutation NameWiseMutation) optional.Of[Node] {
	return mutation.Active(c.Name(), NewObject(c, children))
}

func (c Class) MutateWithObjMutateFunc(children []Node, mutation NameWiseMutation) optional.Of[Node] {
	newChildren := mutateChildren(children, mutation)
	if newChildren.IsNotEmpty() {
		return optional.Value[Node](NewObject(c, newChildren.Value()))
	}

	return mutation.Normal(c.Name(), NewObject(c, children))
}

func (c Class) AppendChildren(children []Node) Object {
	return NewObject(c, children)
}

func (c Class) Name() string {
	return c.name
}

func (c Class) TopLevelString() string {
	return c.String()
}

func (c Class) String() string {
	return c.Name()
}

func (c Class) MutoString() string {
	return c.String()
}

func NewClass(name string) Class {
	return Class{name: name}
}

func UnsafeNodeToClass(obj Node) Class {
	return obj.(Class)
}

var _ MutableNode = Class{}
