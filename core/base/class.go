package base

import (
	"github.com/SSripilaipong/muto/common/optional"
)

type Class interface {
	Node
	ClassType() ClassType
	Children() []Node
	ActivelyMutateWithObjMutateFunc(params ParamChain) optional.Of[Node]
	MutateWithObjMutateFunc(params ParamChain) optional.Of[Node]
	Name() string
	String() string
	Equals(d Class) bool
}

type ClassType string

const (
	ClassTypeRuleBased ClassType = "RULE_BASED"
	ClassTypeImported  ClassType = "IMPORTED"
)

func IsRuleBasedClass(class Class) bool {
	return class.ClassType() == ClassTypeRuleBased
}

func IsImportedClass(class Class) bool {
	return class.ClassType() == ClassTypeImported
}

func UnsafeNodeToClass(obj Node) Class {
	return obj.(Class)
}
