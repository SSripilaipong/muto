package base

import (
	"fmt"

	"muto/common/optional"
)

type NamedObject struct {
	class        Class
	children     []Node
	isTerminated bool
}

func (NamedObject) NodeType() NodeType { return NodeTypeNamedObject }

func (obj NamedObject) ReplaceChild(i int, n Node) Object {
	obj.children[i] = n
	return obj
}

func (obj NamedObject) ActivelyMutateWithObjMutateFunc(objMutate func(string, NamedObject) optional.Of[Node]) optional.Of[Node] {
	return objMutate(obj.Name(), obj)
}

func (obj NamedObject) MutateWithObjMutateFunc(objMutate func(string, NamedObject) optional.Of[Node]) optional.Of[Node] {
	if obj.IsTerminationConfirmed() {
		return optional.Empty[Node]()
	}

	var x Object = obj
	for i, child := range x.Children() {
		if !child.IsTerminationConfirmed() {
			childObj := UnsafeNodeToObject(child)
			if newChild := childObj.MutateWithObjMutateFunc(objMutate); newChild.IsNotEmpty() {
				return optional.Value[Node](obj.ReplaceChild(i, newChild.Value()))
			}
			x = x.ReplaceChild(i, childObj.ConfirmTermination())
		}
	}

	if IsNamedObjectNode(x) {
		namedObj := UnsafeNodeToNamedObject(x)
		r := objMutate(namedObj.Name(), namedObj)
		return r
	}
	return optional.Value[Node](x)
}

func (obj NamedObject) AppendChildren(children []Node) Object {
	obj.children = append(obj.children, children...)
	return obj
}

func (obj NamedObject) Children() []Node {
	return obj.children
}

func (obj NamedObject) IsTerminationConfirmed() bool {
	return obj.isTerminated
}

func (obj NamedObject) ConfirmTermination() Object {
	obj.isTerminated = true
	return obj
}

func (obj NamedObject) LiftTermination() Object {
	obj.isTerminated = false
	return obj
}

func (obj NamedObject) Name() string {
	return obj.class.Name()
}

func (obj NamedObject) String() string {
	if len(obj.Children()) == 0 {
		return fmt.Sprintf(`%s`, obj.Name())
	}
	return fmt.Sprintf(`%s %s`, obj.Name(), objectChildrenToString(obj))
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

func UnsafeObjectToNamedObject(x Object) NamedObject {
	return x.(NamedObject)
}

var _ Object = NamedObject{}
