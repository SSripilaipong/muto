package base

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/common/slc"
)

type CompoundObject struct {
	class      Node
	paramChain ParamChain
}

func (CompoundObject) NodeType() NodeType { return NodeTypeObject }

func (CompoundObject) ObjectType() ObjectType { return ObjectTypeCompound }

func (obj CompoundObject) Mutate() optional.Of[Node] {
	return obj.Head().MutateAsHead(obj.ParamChain())
}

func (obj CompoundObject) ParamChain() ParamChain {
	return obj.paramChain
}

func (obj CompoundObject) MutateAsHead(ParamChain) optional.Of[Node] {
	panic("deprecated")
}

func (obj CompoundObject) AppendChildren(children []Node) Object {
	obj.paramChain = obj.paramChain.AppendChildrenMostOuter(children)
	return obj
}

func (obj CompoundObject) ChainParams(params ParamChain) Object {
	obj.paramChain = obj.paramChain.Chain(params)
	return obj
}

func (obj CompoundObject) Head() Node {
	return obj.class
}

func (obj CompoundObject) Equals(x Object) bool {
	if !NodeEqual(obj.Head(), x.Head()) {
		return false
	}
	if obj.ParamChain().TotalNodes()+x.ParamChain().TotalNodes() == 0 {
		return true
	}
	return objectChildrenEqual(obj.ParamChain().DirectParams(), x.ParamChain().DirectParams())
}

func (obj CompoundObject) String() string {
	return obj.TopLevelString()
}

func (obj CompoundObject) TopLevelString() string {
	return fmt.Sprintf("(%s)", obj.SimplifiedString())
}

func (obj CompoundObject) SimplifiedString() string {
	params := obj.ParamChain()

	if params.Size() == 0 {
		return fmt.Sprintf(`%s`, obj.Head())
	}

	var head Node
	var rest string
	if params.Size() == 1 {
		head = obj.Head()
		rest = objectChildrenToString(params.DirectParams())
	} else {
		head = NewCompoundObject(obj.Head(), params.WithoutMostOuter())
		rest = objectChildrenToString(params.MostOuter())
	}

	if rest == "" {
		return fmt.Sprintf(`%s`, head)
	}
	return fmt.Sprintf(`%s %s`, head, rest)
}

func (obj CompoundObject) MutoString() string {
	return obj.String()
}

func (obj CompoundObject) AppendParams(params ParamChain) Object {
	return NewCompoundObject(obj.Head(), obj.paramChain.AppendAll(params))
}

func NewCompoundObject(class Node, params ParamChain) Object {
	if IsObjectNode(class) {
		panic("fuck") // TODO remove this
	}
	return newCompoundObject(class, params)
}

func newCompoundObject(class Node, params ParamChain) Object {
	return CompoundObject{class: class, paramChain: params}
}

func NewOneLayerObject(class Node, children ...Node) Object {
	return NewCompoundObject(class, NewParamChain(slc.Pure(children)))
}

func NewNamedOneLayerObject(name string, children ...Node) Object {
	return NewOneLayerObject(NewUnlinkedRuleBasedClass(name), children...)
}

func objectChildrenToString(children []Node) string {
	var result []string
	for _, child := range children {
		result = append(result, fmt.Sprint(child))
	}
	return strings.Join(result, " ")
}

func UnsafeNodeToObject(x Node) Object {
	return x.(Object)
}

var _ Object = CompoundObject{}
