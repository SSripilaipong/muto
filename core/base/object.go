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
	return obj.Head().MutateAsHead(obj.Children(), mutation)
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
	head := obj.Head()
	if IsNamedClassNode(head) || IsNumberNode(head) || IsStringNode(head) || IsBooleanNode(head) {
		if len(obj.Children()) == 0 {
			return fmt.Sprintf(`%s`, head)
		}
		return fmt.Sprintf(`%s %s`, head, objectChildrenToString(obj))
	}
	if len(obj.Children()) == 0 {
		return fmt.Sprintf(`(%s)`, head)
	}
	return fmt.Sprintf(`(%s) %s`, head, objectChildrenToString(obj))
}

func NewObject(class Node, children []Node) Object {
	return Object{class: class, children: children}
}

func NewNamedObject(name string, children []Node) Object {
	return Object{class: NewNamedClass(name), children: children}
}

func ObjectToNode(x Object) Node {
	return x
}

func objectChildrenToString(obj Object) string {
	var children []string
	for _, child := range obj.Children() {
		var s string
		if IsObjectNode(child) {
			if len(UnsafeNodeToObject(child).Children()) > 0 {
				s = fmt.Sprintf("(%s)", child)
			} else {
				s = fmt.Sprintf("%s", child)
			}
		} else {
			s = fmt.Sprint(child)
		}
		children = append(children, s)
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
