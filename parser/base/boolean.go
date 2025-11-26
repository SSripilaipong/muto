package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/syntaxtree"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Boolean = ps.Map(
	syntaxtree.NewBoolean,
	ps.First(ps.ToParser(FixedChars("true")), ps.ToParser(FixedChars("false"))),
).Legacy

var BooleanPattern = ps.Map(st.ToPattern, ps.ToParser(Boolean)).Legacy

var BooleanResultNode = ps.Map(booleanToResultNode, ps.ToParser(Boolean)).Legacy

func booleanToResultNode(x syntaxtree.Boolean) stResult.Node { return x }
