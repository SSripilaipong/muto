package base

import (
	"fmt"

	"muto/common/optional"
)

type AnonymousObject struct {
	head         Node
	children     []Node
	isTerminated bool
}

func (AnonymousObject) NodeType() NodeType { return NodeTypeAnonymousObject }

func (obj AnonymousObject) ReplaceChild(i int, n Node) Object {
	obj.children[i] = n
	return obj
}

func (obj AnonymousObject) Children() []Node {
	return obj.children
}

func (obj AnonymousObject) IsTerminationConfirmed() bool {
	return obj.isTerminated
}

func (obj AnonymousObject) ConfirmTermination() Object {
	obj.isTerminated = true
	return obj
}

func (obj AnonymousObject) LiftTermination() Object {
	obj.isTerminated = false
	return obj
}

func (obj AnonymousObject) MutateWithObjMutateFunc(objMutate func(string, NamedObject) optional.Of[Node]) optional.Of[Node] {
	head := obj.Head()
	switch {
	case head.IsTerminationConfirmed():
		return optional.Value[Node](obj.bubbleUp())
	case IsObjectNode(head):
		headObj := UnsafeNodeToObject(head)
		newHead, ok := headObj.MutateWithObjMutateFunc(objMutate).Return()
		if !ok {
			return obj.ReplaceHead(headObj.ConfirmTermination()).MutateWithObjMutateFunc(objMutate)
		}
		return optional.Value[Node](obj.ReplaceHead(newHead))
	}
	panic("not implemented")
}

func (obj AnonymousObject) bubbleUp() Object {
	head := obj.Head()
	if !IsObjectNode(head) {
		return NewDataObject(append([]Node{head}, obj.Children()...))
	}
	return UnsafeNodeToObject(obj.Head()).
		AppendChildren(obj.Children()).
		LiftTermination()
}

func (obj AnonymousObject) AppendChildren(children []Node) Object {
	obj.children = append(obj.children, children...)
	return obj
}

func (obj AnonymousObject) Head() Node {
	return obj.head
}

func (obj AnonymousObject) ReplaceHead(node Node) AnonymousObject {
	obj.head = node
	return obj
}

func (obj AnonymousObject) String() string {
	if len(obj.Children()) == 0 {
		return fmt.Sprintf(`(%s)`, obj.Head())
	}
	return fmt.Sprintf(`(%s) %s`, obj.Head(), objectChildrenToString(obj))
}

func NewAnonymousObject(head Node, children []Node) AnonymousObject {
	return AnonymousObject{head: head, children: children}
}

func AnonymousObjectToNode(obj AnonymousObject) Node {
	return obj
}

func UnsafeObjectToAnonymousObject(obj Object) AnonymousObject {
	return obj.(AnonymousObject)
}

var _ Object = AnonymousObject{}