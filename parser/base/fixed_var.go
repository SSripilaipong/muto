package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var FixedVar = ps.Map(st.NewVariable, ps.Lookahead(not3Dots, identifierStartingWithUpperCase))

var FixedVarPattern = ps.Map(st.ToPattern, FixedVar)

var FixedVarResultNode = ps.Map(stResult.ToNode, FixedVar)

var not3Dots = fn.Not(ps.Matches(ThreeDots))
