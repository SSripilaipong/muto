package base

import (
	"github.com/SSripilaipong/go-common/optional"
)

type Class interface {
	Node
	LinkRule(mutator Rule)
	UnlinkRule()
	ClassType() ClassType
	Children() []Node
	ActivelyMutateWithObjMutateFunc(params ParamChain) optional.Of[Node]
	MutateWithObjMutateFunc(params ParamChain) optional.Of[Node]
	Name() string
	String() string
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

func UnsafeNodeToClass(node Node) Class {
	return node.(Class)
}
