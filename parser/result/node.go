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
			ps.ToParser(EOL(psBase.InParenthesesWhiteSpaceAllowed(nakedObjectMultilines))),
		),
		ps.Map(
			castObjectNodeNaked,
			ps.ToParser(nakedObjectWithChildren()),
		),
		ps.Map(
			wrapNodeWithNakedObject,
			ps.First(
				ps.ToParser(nonNestedNode),
				ps.ToParser(nonObjectNestedNode()),
			),
		),
	).Legacy
}

func anyNode() func([]psBase.Character) tuple.Of2[rslt.Of[stResult.Node], []psBase.Character] {
	return ps.First(
		ps.ToParser(nonNestedNode),
		ps.ToParser(nestedNode()),
	).Legacy
}

var nonNestedNode = ps.First(
	ps.ToParser(psBase.BooleanResultNode),
	ps.ToParser(psBase.StringResultNode),
	ps.ToParser(psBase.RuneResultNode),
	ps.ToParser(psBase.NumberResultNode),
	ps.ToParser(psBase.ClassResultNode),
	ps.ToParser(psBase.TagResultNode),
	ps.ToParser(psBase.FixedVarResultNode),
).Legacy

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
