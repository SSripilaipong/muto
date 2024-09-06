package base

import "muto/common/optional"

type Object interface {
	ClassName() string
	Children() []Node
	NodeType() NodeType
	IsTerminationConfirmed() bool
	ConfirmTermination() Object
	ReplaceChild(i int, n Node) Object
	MutateWithObjMutateFunc(objMutate func(Object) optional.Of[Node]) optional.Of[Node]
}

func UnsafeNodeToObject(x Node) Object {
	return x.(Object)
}
