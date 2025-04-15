package result

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func nakedObjectWithChildren(xs []psBase.Character) []tuple.Of2[objectNode, []psBase.Character] {
	return ps.Or(
		ps.Map(mergeObject, psBase.SpaceSeparated2(objectHead, objectParamPart)),
		ps.Map(fn.Compose(nodeToObject, castObjectNode), psBase.InParenthesesWhiteSpaceAllowed(NakedObjectMultilines)),
	)(xs)
}

var ParseNakedObjectMultilines = fn.Compose3(psBase.FilterStatement, RsNakedObjectMultilines, psBase.StringToCharTokens)

var RsNakedObjectMultilines = ps.Map(fn.Compose(rslt.Value, castObjectNode), NakedObjectMultilines)

func NakedObjectMultilines(xs []psBase.Character) []tuple.Of2[objectNode, []psBase.Character] {
	return ps.Or(
		ps.Map(mergeObject, psBase.WhiteSpaceSeparated2(objectHead, objectParamPartMultilines)),
		ps.Map(nodeToObject, objectHead),
	)(xs)
}

var mergeObject = tuple.Fn2(func(head stResult.Node, params stResult.ParamPart) objectNode {
	return objectNode{head: head, paramPart: params}
})

var nodeToObject = func(head stResult.Node) objectNode {
	return objectNode{head: head, paramPart: stResult.ParamsToFixedParamPart([]stResult.Param{})}
}

func objectHead(xs []psBase.Character) []tuple.Of2[stResult.Node, []psBase.Character] {
	return ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, psBase.InParenthesesWhiteSpaceAllowed(NakedObjectMultilines)),
		ps.Map(stResult.ToNode, structure),
	)(xs)
}
