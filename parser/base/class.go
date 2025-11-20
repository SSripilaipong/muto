package base

import (
	"slices"

	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var DeterminantClass = ps.Map(stBase.ToDeterminant, ps.Filter(validDeterminantClass, Class))

var NonDeterminantClassRulePattern = ps.First(
	ps.Map(stBase.ToPattern, Class),
	ps.Map(stBase.ToPattern, ImportedClass),
)

var ClassResultNode = ps.First(
	ps.Map(stResult.ToNode, Class),
	ps.Map(stResult.ToNode, ImportedClass),
)

var ImportedClass = ps.Map(parseImportedClass, ps.Sequence3(
	ImportPathToken, Dot, Class,
))

var Class = ps.Map(st.NewLocalClass, ps.First(
	ps.Filter(validClassName, identifierStartingWithNonUpperCase),
	ps.Filter(classSymbol, symbol),
))

var parseImportedClass = tuple.Fn3(func(module string, _ Character, class st.LocalClass) st.ImportedClass {
	return st.NewImportedClass(module, class.Name())
})

func validDeterminantClass(class st.LocalClass) bool {
	return !slices.Contains([]string{"try"}, class.Name())
}

func validClassName(x string) bool {
	return !IsBooleanValue(x)
}

func classSymbol(x string) bool {
	return x != "=" && x[0] != '.' && x[0] != '"' && x[0] != '\'' && x[0] != '\\' && x[0] != '('
}
