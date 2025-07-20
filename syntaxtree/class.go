package syntaxtree

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Class interface {
	stResult.Node
	base.Pattern
	base.Determinant

	Name() string
	ClassType() ClassType
}

type ClassType string

const (
	ClassTypeLocal    ClassType = "LOCAL"
	ClassTypeImported ClassType = "IMPORTED"
)

func IsClassTypeLocal(x Class) bool    { return x.ClassType() == ClassTypeLocal }
func IsClassTypeImported(x Class) bool { return x.ClassType() == ClassTypeImported }

func UnsafeRuleResultToClass(p stResult.Node) Class { return p.(Class) }

func UnsafePatternToClass(p base.Pattern) Class { return p.(Class) }
