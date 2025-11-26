package result

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func nakedObjectWithChildren() func(xs []psBase.Character) tuple.Of2[rslt.Of[objectNode], []psBase.Character] {
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
	paramLinesInBlockStartingWithLineBreak := ps.Map(
		mergeParamIgnoreNewLine,
		ps.Sequence2(
			ps.ToParser(psBase.LineBreak),
			ps.ToParser(paramLinesInBlock),
		),
	).Legacy
	paramLineFollowedByParamBlock := ps.Map(
		mergeParamWithFirstParam,
		ps.Sequence3(
			ps.ToParser(paramLine),
			ps.ToParser(psBase.LineBreak),
			ps.ToParser(paramLinesInBlock),
		),
	).Legacy

	head := anyNode()
	return ps.First(
		ps.Map(mergeObject, ps.Sequence2(
			ps.ToParser(head),
			ps.ToParser(paramLinesInBlockStartingWithLineBreak),
		)),
		ps.Map(
			mergeObject,
			ps.ToParser(psBase.SpaceSeparated2(
				head,
				paramLineFollowedByParamBlock,
			)),
		),
		ps.Map(
			mergeObject,
			ps.ToParser(psBase.SpaceSeparated2(
				head,
				paramLine,
			)),
		),
	).Legacy
}

var ParseNakedObjectMultilines = fn.Compose3(psBase.FilterResult, NakedObjectMultilines, psBase.StringToCharTokens)

func NakedObjectMultilines(xs []psBase.Character) tuple.Of2[rslt.Of[stResult.Node], []psBase.Character] {
	return ps.Map(castObjectNode, ps.ToParser(nakedObjectMultilines)).Legacy(xs)
}

func nakedObjectMultilines(xs []psBase.Character) tuple.Of2[rslt.Of[objectNode], []psBase.Character] {
	head := anyNode()
	params := objectParamPartMultilines()
	return ps.First(
		ps.Map(mergeObject, ps.ToParser(psBase.WhiteSpaceSeparated2(head, params))),
		ps.Map(nodeToObject, ps.ToParser(head)),
	).Legacy(xs)
}

var mergeObject = tuple.Fn2(func(head stResult.Node, params stResult.FixedParamPart) objectNode {
	return objectNode{head: head, paramPart: params}
})

var nodeToObject = func(head stResult.Node) objectNode {
	return objectNode{head: head, paramPart: stResult.ParamsToFixedParamPart([]stResult.Param{})}
}
