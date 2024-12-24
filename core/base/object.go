package base

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
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

func (obj Object) Mutate(mutation NameWiseMutation) optional.Of[Node] {
	r, ok := obj.Head().MutateAsHead(obj.ParamChain(), mutation).Return()
	if ok {
		return optional.New(r, ok)
	}
	return obj.BubbleUp()
}

func (obj Object) ParamChain() ParamChain {
	return NewParamChain(slc.Pure(obj.Children()))
}

func (obj Object) MutateAsHead(params ParamChain, mutation NameWiseMutation) optional.Of[Node] {
	newHead, ok := obj.Mutate(mutation).Return()
	if ok {
		return optional.Value[Node](NewObject(newHead, params))
	}
	return optional.Value[Node](obj.AppendChildren(params.DirectParams()))
}

func (obj Object) AppendChildren(children []Node) Object {
	obj.children = append(obj.children, children...)
	return obj
}

func (obj Object) Children() []Node {
	return obj.children
}

func (obj Object) Explode() []Node {
	return append([]Node{obj.Head()}, obj.Children()...)
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

func (obj Object) BubbleUp() optional.Of[Node] {
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

func NewObject(class Node, params ParamChain) Object {
	return Object{class: class, children: params.DirectParams()}
}

func NewOneLayerObject(class Node, children []Node) Object {
	return NewObject(class, NewParamChain(slc.Pure(children)))
}

func NewNamedOneLayerObject(name string, children []Node) Object {
	return NewOneLayerObject(NewClass(name), children)
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

func ExplodeObject(obj Object) []Node {
	return obj.Explode()
}

func UnsafeNodeToObject(x Node) Object {
	return x.(Object)
}

var _ MutableNode = Object{}
