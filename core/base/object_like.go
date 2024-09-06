package base

import "muto/common/optional"

type ObjectLike interface {
	ClassName() string
	Children() []Node
	NodeType() NodeType
	IsTerminationConfirmed() bool
	ConfirmTermination() ObjectLike
	ReplaceChild(i int, n Node) ObjectLike
	MutateWithObjMutateFunc(objMutate func(ObjectLike) optional.Of[Node]) optional.Of[Node]
}

func UnsafeNodeToObjectLike(x Node) ObjectLike {
	return x.(ObjectLike)
}
