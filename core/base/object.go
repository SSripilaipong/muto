package base

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
)

type Object struct {
	class    Node
	children []Node
}

func (Object) NodeType() NodeType { return NodeTypeObject }

func (obj Object) ReplaceChild(i int, n Node) Object {
	obj.children[i] = n
	return obj
}

func (obj Object) Mutate(mutation Mutation) optional.Of[Node] {
	r, ok := obj.Head().MutateAsHead(obj.Children(), mutation).Return()
	if ok {
		return optional.New(r, ok)
	}
	return obj.tryBubbleUp()
}

func (obj Object) MutateAsHead(children []Node, mutation Mutation) optional.Of[Node] {
	newHead, ok := obj.Mutate(mutation).Return()
	if ok {
		return optional.Value[Node](NewObject(newHead, children))
	}
	return optional.Value[Node](obj.AppendChildren(children))
}

func (obj Object) AppendChildren(children []Node) Object {
	obj.children = append(obj.children, children...)
	return obj
}

func (obj Object) Children() []Node {
	return obj.children
}

func (obj Object) Head() Node {
	return obj.class
}

func (obj Object) Equals(x Object) bool {
	if !NodeEqual(obj.Head(), x.Head()) {
		return false
	}
	if len(obj.Children())+len(x.Children()) == 0 {
		return true
	}
	return objectChildrenEqual(obj.Children(), x.Children())
}

func (obj Object) String() string {
	return fmt.Sprintf("(%s)", obj.TopLevelString())
}

func (obj Object) TopLevelString() string {
	if len(obj.Children()) == 0 {
		return fmt.Sprintf(`%s`, obj.Head())
	}
	return fmt.Sprintf(`%s %s`, obj.Head(), objectChildrenToString(obj))
}

func (obj Object) MutoString() string {
	return obj.String()
}

func (obj Object) tryBubbleUp() optional.Of[Node] {
	if len(obj.Children()) == 0 {
		return optional.Value(obj.Head())
	}
	return optional.Empty[Node]()
}

func NewObject(class Node, children []Node) Object {
	return Object{class: class, children: children}
}

func NewNamedObject(name string, children []Node) Object {
	return Object{class: NewClass(name), children: children}
}

func objectChildrenToString(obj Object) string {
	var children []string
	for _, child := range obj.Children() {
		children = append(children, fmt.Sprint(child))
	}
	return strings.Join(children, " ")
}

func ObjectToChildren(obj Object) []Node {
	return obj.Children()
}

func UnsafeNodeToObject(x Node) Object {
	return x.(Object)
}

var _ MutableNode = Object{}
