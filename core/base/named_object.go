package base

import (
	"fmt"
	"strings"

	"muto/common/optional"
)

type NamedObject struct {
	class        Class
	children     []Node
	isTerminated bool
}

func (NamedObject) NodeType() NodeType { return NodeTypeNamedObject }

func (obj NamedObject) ConfirmTermination() ObjectLike {
	obj.isTerminated = true
	return obj
}

func (obj NamedObject) ReplaceChild(i int, n Node) ObjectLike {
	obj.children[i] = n
	return obj
}

func (obj NamedObject) MutateWithObjMutateFunc(objMutate func(ObjectLike) optional.Of[Node]) optional.Of[Node] {
	if obj.IsTerminationConfirmed() {
		return optional.Empty[Node]()
	}

	var x ObjectLike = obj
	for i, child := range x.Children() {
		if !child.IsTerminationConfirmed() {
			childObj := UnsafeNodeToNamedObject(child)
			if newChild := childObj.MutateWithObjMutateFunc(objMutate); newChild.IsNotEmpty() {
				return optional.Value[Node](obj.ReplaceChild(i, newChild.Value()))
			}
			x = x.ReplaceChild(i, childObj.ConfirmTermination())
		}
	}
	return objMutate(x)
}

func (obj NamedObject) IsTerminationConfirmed() bool {
	return obj.isTerminated
}

func (obj NamedObject) Children() []Node {
	return obj.children
}

func (obj NamedObject) Class() Class {
	return obj.class
}

func (obj NamedObject) ClassName() string {
	return obj.Class().Name()
}

func (obj NamedObject) String() string {
	var children []string
	for _, child := range obj.Children() {
		var s string
		if IsNamedObjectNode(child) {
			s = fmt.Sprintf("(%s)", child)
		} else {
			s = fmt.Sprint(child)
		}
		children = append(children, s)
	}
	jointChildren := strings.Join(children, " ")
	return fmt.Sprintf(`%s %s`, obj.ClassName(), jointChildren)
}

func NewNamedObject(name string, children []Node) NamedObject {
	return NamedObject{class: NewNamedClass(name), children: children}
}

func NamedObjectToNode(x NamedObject) Node {
	return x
}

func UnsafeNodeToNamedObject(x Node) NamedObject {
	return x.(NamedObject)
}
