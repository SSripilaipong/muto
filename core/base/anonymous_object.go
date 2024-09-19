package base

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
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

func (obj AnonymousObject) ConfirmTermination() MutableNode {
	obj.isTerminated = true
	return obj
}

func (obj AnonymousObject) LiftTermination() MutableNode {
	obj.isTerminated = false
	return obj
}

func (obj AnonymousObject) Mutate(mutation Mutation) optional.Of[Node] {
	head := obj.Head()
	switch {
	case head.IsTerminationConfirmed():
		return optional.Value(obj.bubbleUp())
	case IsMutableNode(head):
		headObj := UnsafeNodeToMutable(head)
		newHead, ok := headObj.Mutate(mutation).Return()
		if !ok {
			return obj.ReplaceHead(headObj.ConfirmTermination()).Mutate(mutation)
		}
		return optional.Value[Node](obj.ReplaceHead(newHead))
	}
	panic("not implemented")
}

func (obj AnonymousObject) bubbleUp() Node {
	head := obj.Head()
	switch {
	case !IsMutableNode(head):
		if len(obj.Children()) > 0 {
			return NewDataObject(append([]Node{head}, obj.Children()...))
		}
		return head
	case IsObjectNode(head):
		return UnsafeNodeToObject(head).
			AppendChildren(obj.Children()).
			LiftTermination()
	case IsNamedClassNode(head):
		class := UnsafeNodeToNamedClass(head)
		if len(obj.Children()) > 0 {
			return NewObject(class, obj.Children())
		}
		return head
	}
	panic("not implemented")
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

func UnsafeNodeToAnonymousObject(obj Node) AnonymousObject {
	return obj.(AnonymousObject)
}

var _ Object = AnonymousObject{}
