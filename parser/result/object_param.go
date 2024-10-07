package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func objectParamPart(xs []psBase.Character) []tuple.Of2[stResult.ParamPart, []psBase.Character] {
	return ps.Map(stResult.ParamsToParamPart, psBase.OptionalGreedyRepeatSpaceSeparated(objectParam))(xs)
}

func objectParam(xs []psBase.Character) []tuple.Of2[stResult.Param, []psBase.Character] {
	return ps.Or(
		ps.Map(castNodeToParam, nonNestedNode),
		psBase.VariadicVarResultNode,
		ps.Map(castObjectParam, psBase.InParentheses(object)),
	)(xs)
}

func castNodeToParam(x stResult.Node) stResult.Param {
	return x
}

func castObjectParam(x objectNode) stResult.Param {
	return stResult.NewObject(x.Head(), x.ParamPart())
}
