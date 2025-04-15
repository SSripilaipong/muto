package result

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var ParseNode = fn.Compose3(psBase.FilterStatement, RsNode, psBase.StringToCharTokens)

var RsNode = ps.Map(rslt.Value, Node())

func Node() func([]psBase.Character) []tuple.Of2[stResult.Node, []psBase.Character] {
	return ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, psBase.InParentheses(NakedObjectMultilines)),
		ps.Map(stResult.ToNode, structure),
	)
}

var RsSimplifiedNode = ps.Map(rslt.Value, SimplifiedNode())

func SimplifiedNode() func([]psBase.Character) []tuple.Of2[stResult.Node, []psBase.Character] {
	return ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, nakedObjectWithChildren),
		ps.Map(stResult.ToNode, structure),
	)
}

var nonNestedNode = ps.Or(
	psBase.BooleanResultNode,
	psBase.StringResultNode,
	psBase.NumberResultNode,
	psBase.ClassResultNode,
	psBase.TagResultNode,
	psBase.FixedVarResultNode,
)

func castObjectNode(obj objectNode) stResult.Node {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}
