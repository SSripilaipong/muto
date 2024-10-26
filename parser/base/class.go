package base

import (
	"slices"

	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Class = ps.Map(st.NewClass, ps.Or(
	ps.Filter(validClassName, identifierStartingWithLowerCase),
	ps.Filter(classSymbol, symbol),
))

var ClassRule = ps.Filter(validClassRule, Class)

func validClassRule(class st.Class) bool {
	return !slices.Contains([]string{"try"}, class.Name())
}

var ClassResultNode = ps.Map(classToResultNode, Class)

func validClassName(x string) bool {
	return !IsBooleanValue(x)
}

func classSymbol(x string) bool {
	return x != "=" && x[0] != '.' && x[0] != '"'
}

func classToResultNode(x st.Class) stResult.Node { return x }
