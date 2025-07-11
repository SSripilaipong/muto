package base

import (
	"slices"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Class = ps.Map(syntaxtree.NewClass, ps.Or(
	ps.Filter(validClassName, identifierStartingWithNonUpperCase),
	ps.Filter(classSymbol, symbol),
))

var ClassRule = ps.Filter(validClassRule, Class)

var ClassRulePattern = ps.Map(st.ToPattern, ClassRule)
var ClassDeterminant = ps.Map(st.ToDeterminant, ClassRule)

func validClassRule(class syntaxtree.Class) bool {
	return !slices.Contains([]string{"try"}, class.Name())
}

var ClassResultNode = ps.Map(classToResultNode, Class)

func validClassName(x string) bool {
	return !IsBooleanValue(x)
}

func classSymbol(x string) bool {
	return x != "=" && x[0] != '.' && x[0] != '"'
}

func classToResultNode(x syntaxtree.Class) stResult.Node { return x }
