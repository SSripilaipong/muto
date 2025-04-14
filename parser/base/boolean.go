package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Boolean = ps.Map(st.NewBoolean, ps.Or(fixedChars("true"), fixedChars("false")))

var BooleanPattern = ps.Map(st.ToPattern, Boolean)

var BooleanResultNode = ps.Map(booleanToResultNode, Boolean)

func booleanToResultNode(x st.Boolean) stResult.Node { return x }
