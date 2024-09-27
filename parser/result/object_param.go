package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func objectParamPart(xs []tokenizer.Token) []tuple.Of2[stResult.ParamPart, []tokenizer.Token] {
	return ps.Map(stResult.ParamsToParamPart, ps.OptionalGreedyRepeat(objectParam))(xs)
}

func objectParam(xs []tokenizer.Token) []tuple.Of2[stResult.Param, []tokenizer.Token] {
	return ps.Or(
		ps.Map(castNodeToParam, nonNestedNode),
		ps.Map(castVariadicVariableParam, psBase.VariadicVar),
		ps.Map(castObjectParam, psBase.InParentheses(object)),
	)(xs)
}

func castNodeToParam(x stResult.Node) stResult.Param {
	return x
}

func castObjectParam(x objectNode) stResult.Param {
	return stResult.NewObject(x.Head(), x.ParamPart())
}

func castVariadicVariableParam(x psBase.VariadicVarNode) stResult.Param {
	return stResult.NewVariadicVariable(x.Name())
}
