package base

import (
	"github.com/SSripilaipong/muto/common/optional"
)

type Mutator interface {
	Active(obj Object) optional.Of[Node]
	Normal(obj Object) optional.Of[Node]
}

type Class struct {
	name    string
	mutator Mutator
}

func (c *Class) Children() []Node {
	return nil
}

func (c *Class) NodeType() NodeType {
	return NodeTypeClass
}

func (c *Class) MutateAsHead(params ParamChain) optional.Of[Node] {
	if result, ok := c.ActivelyMutateWithObjMutateFunc(params).Return(); ok {
		return optional.Value(result)
	}
	return c.MutateWithObjMutateFunc(params)
}

func (c *Class) ActivelyMutateWithObjMutateFunc(params ParamChain) optional.Of[Node] {
	if c.mutator == nil {
		return optional.Empty[Node]()
	}
	return c.mutator.Active(NewCompoundObject(c, params))
}

func (c *Class) MutateWithObjMutateFunc(params ParamChain) optional.Of[Node] {
	newChildren := MutateParamChain(params)
	if newChildren.IsNotEmpty() {
		return optional.Value[Node](NewCompoundObject(c, newChildren.Value()))
	}
	if c.mutator == nil {
		return optional.Empty[Node]()
	}
	return c.mutator.Normal(NewCompoundObject(c, params))
}

func (c *Class) Link(mutator Mutator) {
	c.mutator = mutator
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) TopLevelString() string {
	return c.String()
}

func (c *Class) String() string {
	return c.Name()
}

func (c *Class) MutoString() string {
	return c.String()
}

func (c *Class) Equals(d *Class) bool { return c.Name() == d.Name() }

func NewClass(name string, mutator Mutator) *Class {
	return &Class{name: name, mutator: mutator}
}

func NewUnlinkedClass(name string) *Class {
	return NewClass(name, nil)
}

func UnsafeNodeToClass(obj Node) *Class {
	return obj.(*Class)
}

var _ Node = &Class{}
