package base

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
)

type Object interface {
	Children() []Node
	NodeType() NodeType
	IsTerminationConfirmed() bool
	ConfirmTermination() MutableNode
	LiftTermination() MutableNode
	ReplaceChild(i int, n Node) Object
	Mutate(mutation Mutation) optional.Of[Node]
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

var _ MutableNode = Object(nil)
