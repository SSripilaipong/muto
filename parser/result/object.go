package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func object(xs []psBase.Character) []tuple.Of2[objectNode, []psBase.Character] {
	return ps.Or(
		ps.Map(mergeObject, psBase.SpaceSeparated2(objectHead, objectParamPart)),
		ps.Map(objectHeadToObject, objectHead),
	)(xs)
}

func objectMultilines(xs []psBase.Character) []tuple.Of2[objectNode, []psBase.Character] {
	return ps.Or(
		ps.Map(mergeObject, psBase.WhiteSpaceSeparated2(objectHead, objectParamPartMultilines)),
		ps.Map(objectHeadToObject, objectHead),
	)(xs)
}

var mergeObject = tuple.Fn2(func(head stResult.Node, params stResult.ParamPart) objectNode {
	return objectNode{head: head, paramPart: params}
})

var objectHeadToObject = func(head stResult.Node) objectNode {
	return objectNode{head: head, paramPart: stResult.ParamsToFixedParamPart([]stResult.Param{})}
}

func objectHead(xs []psBase.Character) []tuple.Of2[stResult.Node, []psBase.Character] {
	return ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, psBase.InParenthesesWhiteSpaceAllowed(objectMultilines)),
		ps.Map(stResult.ToNode, structure),
	)(xs)
}
