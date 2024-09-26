package base

import "github.com/SSripilaipong/muto/common/optional"

type MutableNode interface {
	Node
	Mutate(mutation Mutation) optional.Of[Node]
}

func UnsafeNodeToMutable(x Node) MutableNode {
	return x.(MutableNode)
}
