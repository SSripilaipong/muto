package base

import (
	"slices"

	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var DeterminantClass = ps.Map(stBase.ToDeterminant, ps.Filter(validDeterminantClass, ps.ToParser(Class))).Legacy

var NonDeterminantClassRulePattern = ps.First(
	ps.Map(stBase.ToPattern, ps.ToParser(ImportedClass)),
	ps.Map(stBase.ToPattern, ps.ToParser(Class)),
).Legacy

var ClassResultNode = ps.First(
	ps.Map(stResult.ToNode, ps.ToParser(ImportedClass)),
	ps.Map(stResult.ToNode, ps.ToParser(Class)),
).Legacy

var ImportedClass = ps.Map(parseImportedClass, ps.Sequence3(
	ps.ToParser(ImportPathToken), ps.ToParser(Dot), ps.ToParser(Class),
)).Legacy

var Class = ps.Map(st.NewLocalClass, ps.First(
	ps.Filter(validClassName, identifierStartingWithNonUpperCase),
	ps.Filter(classSymbol, symbol),
)).Legacy

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
