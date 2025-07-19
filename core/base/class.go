package base

import (
	"github.com/SSripilaipong/muto/common/optional"
)

type Class interface {
	Node
	Children() []Node
	ActivelyMutateWithObjMutateFunc(params ParamChain) optional.Of[Node]
	MutateWithObjMutateFunc(params ParamChain) optional.Of[Node]
	Name() string
	String() string
	Equals(d Class) bool
}

func UnsafeNodeToClass(obj Node) Class {
	return obj.(Class)
}
