package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var FixedVar = ps.Map(st.NewVariable, ps.Lookahead(not3Dots, identifierStartingWithUpperCase))

var FixedVarPatternParam = ps.Map(fixedVarToPatternParam, FixedVar)

var FixedVarResultNode = ps.Map(fixedVarToResultNode, FixedVar)

func fixedVarToPatternParam(x st.Variable) st.PatternParam { return x }
func fixedVarToResultNode(x st.Variable) stResult.Node     { return x }

var not3Dots = fn.Not(ps.Matches(ThreeDots))
