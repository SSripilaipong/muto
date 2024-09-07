package base

import (
	"fmt"
	"strings"

	"muto/common/optional"
)

type Object interface {
	Children() []Node
	NodeType() NodeType
	IsTerminationConfirmed() bool
	ConfirmTermination() Object
	LiftTermination() Object
	ReplaceChild(i int, n Node) Object
	MutateWithObjMutateFunc(objMutate func(string, NamedObject) optional.Of[Node]) optional.Of[Node]
	AppendChildren(children []Node) Object
}

func objectChildrenToString(obj Object) string {
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
	return strings.Join(children, " ")
}

func UnsafeNodeToObject(x Node) Object {
	return x.(Object)
}
