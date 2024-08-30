package base

import (
	"fmt"
	"strings"
)

type Object struct {
	class        Class
	children     []Node
	isTerminated bool
}

func (Object) NodeType() NodeType { return NodeTypeObject }

func (obj Object) Terminate() ObjectLike {
	obj.isTerminated = true
	return obj
}

func (obj Object) IsTerminated() bool {
	return obj.isTerminated
}

func (obj Object) Children() []Node {
	return obj.children
}

func (obj Object) Class() Class {
	return obj.class
}

func (obj Object) ClassName() string {
	return obj.Class().Name()
}

func (obj Object) String() string {
	var children []string
	for _, child := range obj.Children() {
		children = append(children, fmt.Sprint(child))
	}
	jointChildren := strings.Join(children, " ")
	return fmt.Sprintf(`%s %s`, obj.ClassName(), jointChildren)
}

func NewObject(class Class, children []Node) Object {
	return Object{class: class, children: children}
}

func ObjectToNode(x Object) Node {
	return x
}

func UnsafeNodeToObject(x Node) Object {
	return x.(Object)
}
