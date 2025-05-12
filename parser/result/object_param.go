package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func objectParamPart() func(xs []psBase.Character) []tuple.Of2[stResult.FixedParamPart, []psBase.Character] {
	return ps.Map(stResult.ParamsToFixedParamPart, psBase.OptionalGreedyRepeatSpaceSeparated(objectParam))
}

func objectParamPartMultilines(xs []psBase.Character) []tuple.Of2[stResult.FixedParamPart, []psBase.Character] {
	return ps.Map(stResult.ParamsToFixedParamPart, psBase.OptionalGreedyRepeatWhiteSpaceSeparated(objectParam))(xs)
}

func objectParam(xs []psBase.Character) []tuple.Of2[stResult.Param, []psBase.Character] {
	return ps.Or(
		ps.Map(stResult.ToParam, ps.Or(nonNestedNode, nestedNode())),
		psBase.VariadicVarResultNode,
	)(xs)
}

func castObjectParam(x objectNode) stResult.Param {
	return stResult.NewObject(x.Head(), x.ParamPart())
}
