package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Class = ps.Map(st.NewClass, ps.Or(
	ps.Filter(validClassName, identifierStartingWithLowerCase),
	ps.Filter(classSymbol, symbol),
))

var ClassResultNode = ps.Map(classToResultNode, Class)

func validClassName(x string) bool {
	return !IsBooleanValue(x)
}

func classSymbol(x string) bool {
	return x != "="
}

func classToResultNode(x st.Class) stResult.Node { return x }
