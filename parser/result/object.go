package result

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func nakedObjectWithChildren() func(xs []psBase.Character) []tuple.Of2[objectNode, []psBase.Character] {
	mergeParam := func(ps []stResult.FixedParamPart) (r stResult.FixedParamPart) {
		for _, p := range ps {
			r = r.Append(p)
		}
		return
	}
	mergeParamIgnoreNewLine := tuple.Fn2(func(_ psBase.Character, ps []stResult.FixedParamPart) stResult.FixedParamPart {
		return mergeParam(ps)
	})
	mergeParamWithFirstParam := tuple.Fn3(func(first stResult.FixedParamPart, _ psBase.Character, others []stResult.FixedParamPart) stResult.FixedParamPart {
		return mergeParam(append([]stResult.FixedParamPart{first}, others...))
	})

	paramLine := objectParamPart()
	paramLinesInBlock := psBase.GreedyRepeatedLinesInAutoIndentBlockAtLeastOnce(paramLine)
	paramLinesInBlockStartingWithLineBreak := ps.Map(mergeParamIgnoreNewLine, ps.Sequence2(psBase.LineBreak, paramLinesInBlock))
	paramLineFollowedByParamBlock := ps.Map(mergeParamWithFirstParam, ps.Sequence3(paramLine, psBase.LineBreak, paramLinesInBlock))

	return ps.Or(
		ps.Map(mergeObject, ps.Sequence2(objectHead, paramLinesInBlockStartingWithLineBreak)),
		ps.Map(mergeObject, psBase.SpaceSeparated2(objectHead, paramLineFollowedByParamBlock)),
		ps.Map(mergeObject, psBase.SpaceSeparated2(objectHead, paramLine)),
	)
}

var ParseNakedObjectMultilines = fn.Compose3(psBase.FilterStatement, RsNakedObjectMultilines, psBase.StringToCharTokens)

var RsNakedObjectMultilines = ps.Map(fn.Compose(rslt.Value, castObjectNode), NakedObjectMultilines)

func NakedObjectMultilines(xs []psBase.Character) []tuple.Of2[objectNode, []psBase.Character] {
	return ps.Or(
		ps.Map(mergeObject, psBase.WhiteSpaceSeparated2(objectHead, objectParamPartMultilines)),
		ps.Map(nodeToObject, objectHead),
	)(xs)
}

var mergeObject = tuple.Fn2(func(head stResult.Node, params stResult.FixedParamPart) objectNode {
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
