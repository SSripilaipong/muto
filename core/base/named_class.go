package base

import (
	"github.com/SSripilaipong/muto/common/optional"
)

type NamedClass struct {
	name         string
	isTerminated bool
}

func (c NamedClass) Children() []Node {
	return nil
}

func (c NamedClass) NodeType() NodeType {
	return NodeTypeNamedClass
}

func (c NamedClass) IsTerminationConfirmed() bool {
	return c.isTerminated
}

func (c NamedClass) ConfirmTermination() MutableNode {
	c.isTerminated = true
	return c
}

func (c NamedClass) LiftTermination() MutableNode {
	c.isTerminated = false
	return c
}

func (c NamedClass) Mutate(mutation Mutation) optional.Of[Node] {
	return NewObject(c, nil).Mutate(mutation)
}

func (c NamedClass) AppendChildren(children []Node) Object {
	return NewObject(c, children)
}

func (c NamedClass) Name() string {
	return c.name
}

func (c NamedClass) String() string {
	return c.Name()
}

func NewNamedClass(name string) NamedClass {
	return NamedClass{name: name}
}

func UnsafeNodeToNamedClass(obj Node) NamedClass {
	return obj.(NamedClass)
}

var _ Class = NamedClass{}
