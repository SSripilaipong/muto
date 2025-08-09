package result

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var ParseSimplifiedNode = fn.Compose3(psBase.FilterStatement, RsSimplifiedNode, psBase.StringToCharTokens)

var RsSimplifiedNode = ps.Map(rslt.Value, SimplifiedNode())

func SimplifiedNode() func([]psBase.Character) []tuple.Of2[stResult.SimplifiedNode, []psBase.Character] {
	return ps.Or(
		ps.Map(castObjectNodeSimplified, psBase.InParenthesesWhiteSpaceAllowed(nakedObjectMultilines)),
		ps.Map(castObjectNodeNaked, nakedObjectWithChildren()),
		ps.Map(wrapNodeWithNakedObject, ps.Or(nonNestedNode, nonObjectNestedNode())),
	)
}

func anyNode() func([]psBase.Character) []tuple.Of2[stResult.Node, []psBase.Character] {
	return ps.Or(nonNestedNode, nestedNode())
}

var nonNestedNode = ps.Or(
	psBase.BooleanResultNode,
	psBase.StringResultNode,
	psBase.RuneResultNode,
	psBase.NumberResultNode,
	psBase.ClassResultNode,
	psBase.TagResultNode,
	psBase.FixedVarResultNode,
)

func castObjectNode(obj objectNode) stResult.Node {
	return castObjectResult(obj)
}

func castObjectNodeSimplified(obj objectNode) stResult.SimplifiedNode {
	return castObjectResult(obj)
}

func castObjectNodeNaked(obj objectNode) stResult.SimplifiedNode {
	return stResult.NewNakedObject(obj.Head(), obj.ParamPart())
}

func castObjectResult(obj objectNode) stResult.Object {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}

var wrapNodeWithNakedObject = fn.Compose(stResult.ToSimplifiedNode, stResult.SingleNodeToNakedObject)
