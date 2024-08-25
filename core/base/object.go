package base

type Object struct {
	class    Class
	children []Node
}

func (obj Object) NodeType() NodeType {
	return NodeTypeObject
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

func NewObject(class Class, children []Node) Object {
	return Object{class: class, children: children}
}
