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
	Mutate(mutation Mutation) optional.Of[Node]
	ActivelyMutateWithObjMutateFunc(mutation Mutation) optional.Of[Node]
	MutateWithObjMutateFunc(mutation Mutation) optional.Of[Node]
	AppendChildren(children []Node) Object
}

func objectChildrenToString(obj Object) string {
	var children []string
	for _, child := range obj.Children() {
		var s string
		if IsNamedObjectNode(child) {
			if len(UnsafeNodeToNamedObject(child).Children()) > 0 {
				s = fmt.Sprintf("(%s)", child)
			} else {
				s = fmt.Sprintf("%s", child)
			}
		} else if IsAnonymousObjectNode(child) {
			s = fmt.Sprintf("(%s)", child)
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

func EqualObject(x, y Object) bool {
	return x == y
}
