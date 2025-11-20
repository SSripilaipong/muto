package result

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func objectParamPart() func(xs []psBase.Character) tuple.Of2[rslt.Of[stResult.FixedParamPart], []psBase.Character] {
	return ps.Map(
		stResult.ParamsToFixedParamPart,
		psBase.OptionalGreedyRepeatSpaceSeparated(objectParam),
	)
}

func objectParamPartMultilines() func(xs []psBase.Character) tuple.Of2[rslt.Of[stResult.FixedParamPart], []psBase.Character] {
	return ps.Map(
		stResult.ParamsToFixedParamPart,
		psBase.OptionalGreedyRepeatWhiteSpaceSeparated(objectParam),
	)
}

func objectParam(xs []psBase.Character) tuple.Of2[rslt.Of[stResult.Param], []psBase.Character] {
	return ps.First(
		ps.Map(stResult.ToParam, ps.First(
			nonNestedNode,
			nestedNode(),
		)),
		psBase.VariadicVarResultNode,
	)(xs)
}
