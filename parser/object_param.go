package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func objectParamPart(xs []tokenizer.Token) []tuple.Of2[stResult.ParamPart, []tokenizer.Token] {
	return ps.Map(stResult.ParamsToParamPart, ps.OptionalGreedyRepeat(objectParam))(xs)
}

func objectParam(xs []tokenizer.Token) []tuple.Of2[stResult.Param, []tokenizer.Token] {

	castObject := func(obj objectNode) stResult.Param {
		return stResult.NewObject(obj.Head(), obj.ParamPart())
	}
	castClass := func(objName tokenizer.Token) stResult.Param {
		return st.NewClass(objName.Value())
	}
	castBoolean := func(x tokenizer.Token) stResult.Param {
		return st.NewBoolean(x.Value())
	}
	castNumber := func(x tokenizer.Token) stResult.Param {
		return st.NewNumber(x.Value())
	}
	castString := func(x tokenizer.Token) stResult.Param {
		s := x.Value()
		return st.NewString(s[1 : len(s)-1])
	}
	castVariable := func(x tokenizer.Token) stResult.Param {
		return st.NewVariable(x.Value())
	}
	castVariadicVariable := func(x variadicVarNode) stResult.Param {
		return stResult.NewVariadicVariable(x.Name())
	}

	return ps.Or(
		ps.Map(castBoolean, psBase.Boolean),
		ps.Map(castString, psBase.String),
		ps.Map(castNumber, psBase.Number),
		ps.Map(castVariadicVariable, variadicVar),
		ps.Map(castVariable, psBase.Variable),
		ps.Map(castClass, psBase.ClassIncludingSymbols),
		ps.Map(castObject, psBase.InParentheses(object)),
	)(xs)
}
