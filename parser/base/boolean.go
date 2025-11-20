package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Boolean = ps.Map(
	syntaxtree.NewBoolean,
	ps.First(FixedChars("true"), FixedChars("false")),
)

var BooleanPattern = ps.Map(st.ToPattern, Boolean)

var BooleanResultNode = ps.Map(booleanToResultNode, Boolean)

func booleanToResultNode(x syntaxtree.Boolean) stResult.Node { return x }
