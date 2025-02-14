package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
)

type PrimitiveObject struct {
	class Node
}

func (PrimitiveObject) NodeType() NodeType { return NodeTypeObject }

func (PrimitiveObject) ObjectType() ObjectType { return ObjectTypePrimitive }

func (obj PrimitiveObject) Mutate(mutation NameWiseMutation) optional.Of[Node] {
	return obj.Head().MutateAsHead(obj.ParamChain(), mutation)
}

func (obj PrimitiveObject) ParamChain() ParamChain {
	return NewParamChain([][]Node{{}})
}

func (obj PrimitiveObject) MutateAsHead(ParamChain, NameWiseMutation) optional.Of[Node] {
	panic("deprecated")
}

func (obj PrimitiveObject) AppendChildren(children []Node) Object {
	return NewCompoundObject(obj.Head(), obj.ParamChain().AppendChildrenMostOuter(children))
}

func (obj PrimitiveObject) ChainParams(params ParamChain) Object {
	return NewCompoundObject(obj.Head(), obj.ParamChain().Chain(params))
}

func (obj PrimitiveObject) Children() []Node {
	return obj.ParamChain().DirectParams()
}

func (obj PrimitiveObject) Head() Node {
	return obj.class
}

func (obj PrimitiveObject) Equals(x Object) bool {
	if !NodeEqual(obj.Head(), x.Head()) {
		return false
	}
	if obj.ParamChain().TotalNodes()+x.ParamChain().TotalNodes() == 0 {
		return true
	}
	return objectChildrenEqual(obj.Children(), x.Children())
}

func (obj PrimitiveObject) String() string {
	return obj.TopLevelString()
}

func (obj PrimitiveObject) TopLevelString() string {
	return fmt.Sprint(obj.Head())
}

func (obj PrimitiveObject) MutoString() string {
	return obj.String()
}

func (obj PrimitiveObject) BubbleUp() optional.Of[Node] {
	children := obj.Children()

	if len(children) == 0 {
		return optional.Value(obj.Head())
	}

	head := obj.Head()
	if IsObjectNode(head) {
		return optional.Value[Node](UnsafeNodeToObject(head).AppendChildren(children))
	}

	return optional.Empty[Node]()
}

func (obj PrimitiveObject) AppendParams(params ParamChain) Object {
	return NewCompoundObject(obj.Head(), obj.ParamChain().AppendAll(params))
}

func NewPrimitiveObject[T Node](class T) Object {
	return PrimitiveObject{class: class}
}

var _ Object = PrimitiveObject{}
