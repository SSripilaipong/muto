package base

import (
	"github.com/SSripilaipong/muto/common/fn"
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
	return c.MutateAsHead(NewParamChain([][]Node{{}}), mutation) // act as a one-layered object with no param
}

func (c Class) MutateAsHead(params ParamChain, mutation NameWiseMutation) optional.Of[Node] {
	if result, ok := c.ActivelyMutateWithObjMutateFunc(params, mutation).Return(); ok {
		return optional.Value(result)
	}
	return c.MutateWithObjMutateFunc(params, mutation)
}

func (c Class) ActivelyMutateWithObjMutateFunc(params ParamChain, mutation NameWiseMutation) optional.Of[Node] {
	return mutation.Active(c.Name(), NewCompoundObject(c, params))
}

func (c Class) MutateWithObjMutateFunc(params ParamChain, mutation NameWiseMutation) optional.Of[Node] {
	newChildren := mutateParamChain(params, mutation)
	if newChildren.IsNotEmpty() {
		return optional.Value[Node](NewCompoundObject(c, newChildren.Value()))
	}
	return mutation.Normal(c.Name(), NewCompoundObject(c, params))
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

var NewClassObject = fn.Compose(NewPrimitiveObject, NewClass)

func UnsafeNodeToClass(obj Node) Class {
	return obj.(Class)
}

var _ MutableNode = Class{}
