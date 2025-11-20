package result

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var ParseSimplifiedNode = fn.Compose3(psBase.FilterResult, SimplifiedNodeInstant, psBase.StringToCharTokens)

var SimplifiedNodeInstant = SimplifiedNode()

func SimplifiedNode() func([]psBase.Character) tuple.Of2[rslt.Of[stResult.SimplifiedNode], []psBase.Character] {
	return ps.First(
		ps.Map(
			castObjectNodeSimplified,
			EOL(psBase.InParenthesesWhiteSpaceAllowed(nakedObjectMultilines)),
		),
		ps.Map(
			castObjectNodeNaked,
			nakedObjectWithChildren(),
		),
		ps.Map(
			wrapNodeWithNakedObject,
			ps.First(
				nonNestedNode,
				nonObjectNestedNode(),
			),
		),
	)
}

func anyNode() func([]psBase.Character) tuple.Of2[rslt.Of[stResult.Node], []psBase.Character] {
	return ps.First(
		nonNestedNode,
		nestedNode(),
	)
}

var nonNestedNode = ps.First(
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
