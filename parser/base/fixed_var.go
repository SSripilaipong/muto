package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var FixedVarWithUnderscore = ps.Map(syntaxtree.NewVariable, ps.Lookahead(not3Dots, identifierStartingWithUpperCaseAndUnderscore)).Legacy
var FixedVar = ps.Map(
	syntaxtree.NewVariable,
	ps.Lookahead(not3Dots, identifierStartingWithUpperCase),
).Legacy

var FixedVarWithUnderscorePattern = ps.Map(st.ToPattern, ps.ToParser(FixedVarWithUnderscore)).Legacy

var FixedVarResultNode = ps.Map(stResult.ToNode, ps.ToParser(FixedVar)).Legacy

var not3Dots = fn.Not(ps.Matches(ps.ToParser(ThreeDots)))
