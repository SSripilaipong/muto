package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var FixedVarWithUnderscore = ps.Map(st.NewVariable, ps.Lookahead(not3Dots, identifierStartingWithUpperCaseAndUnderscore))
var FixedVar = ps.Map(st.NewVariable, ps.Lookahead(not3Dots, identifierStartingWithUpperCase))

var FixedVarWithUnderscorePattern = ps.Map(st.ToPattern, FixedVarWithUnderscore)

var FixedVarResultNode = ps.Map(stResult.ToNode, FixedVar)

var not3Dots = fn.Not(ps.Matches(ThreeDots))
