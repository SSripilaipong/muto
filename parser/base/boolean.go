package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Boolean = ps.Map(st.NewBoolean, ps.Or(fixedChars("true"), fixedChars("false")))

var BooleanPatternParam = ps.Map(booleanToPatternParam, Boolean)

var BooleanResultNode = ps.Map(booleanToResultNode, Boolean)

func booleanToPatternParam(x st.Boolean) st.PatternParam { return x }
func booleanToResultNode(x st.Boolean) stResult.Node     { return x }
