package base

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
)

type Object struct {
	class      Node
	paramChain ParamChain
}

func (Object) NodeType() NodeType { return NodeTypeObject }

func (obj Object) ReplaceChild(i int, n Node) Object {
	obj.paramChain = obj.paramChain.ReplaceChild(0, i, n).Value() // assume no error TODO revise
	return obj
}

func (obj Object) Mutate(mutation NameWiseMutation) optional.Of[Node] {
	return obj.Head().MutateAsHead(obj.ParamChain(), mutation)
}

func (obj Object) ParamChain() ParamChain {
	return obj.paramChain
}

func (obj Object) MutateAsHead(ParamChain, NameWiseMutation) optional.Of[Node] {
	panic("deprecated")
}

func (obj Object) AppendChildren(children []Node) Object {
	obj.paramChain = obj.paramChain.AppendChildrenMostOuter(children)
	return obj
}

func (obj Object) ChainParams(params ParamChain) Object {
	obj.paramChain = obj.paramChain.Chain(params)
	return obj
}

func (obj Object) Children() []Node {
	return obj.ParamChain().DirectParams()
}

func (obj Object) Head() Node {
	return obj.class
}

func (obj Object) Equals(x Object) bool {
	if !NodeEqual(obj.Head(), x.Head()) {
		return false
	}
	if obj.ParamChain().TotalNodes()+x.ParamChain().TotalNodes() == 0 {
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
	return Object{class: class, paramChain: params}
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

func UnsafeNodeToObject(x Node) Object {
	return x.(Object)
}

var _ MutableNode = Object{}
