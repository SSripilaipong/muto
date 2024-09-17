package base

import "muto/common/optional"

type MutableNode interface {
	NodeType() NodeType
	IsTerminationConfirmed() bool
	ConfirmTermination() MutableNode
	LiftTermination() MutableNode
	Mutate(mutation Mutation) optional.Of[Node]
}

func UnsafeNodeToMutable(x Node) MutableNode {
	return x.(MutableNode)
}
