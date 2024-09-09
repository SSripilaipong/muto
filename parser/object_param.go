package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	st "muto/syntaxtree"
)

func objectParamPart(xs []tokenizer.Token) []tuple.Of2[st.ObjectParamPart, []tokenizer.Token] {
	return ps.Or(
		ps.Map(mergeLeftVariadicParamPart, ps.Sequence2(variadicVar, ps.OptionalGreedyRepeat(objectParam))),
		ps.Map(mergeRightVariadicParamPart, ps.Sequence2(ps.GreedyRepeatAtLeastOnce(objectParam), variadicVar)),
		ps.Map(st.ObjectParamsToObjectParamPart, ps.OptionalGreedyRepeat(objectParam)),
	)(xs)
}

func objectParam(xs []tokenizer.Token) []tuple.Of2[st.ObjectParam, []tokenizer.Token] {
	return ps.Or(
		objectParamValue,
		ps.Map(variableToObjectParam, variable),
		ps.Map(namedObjectWithoutChildToObjectParam, objectName),
		ps.Map(namedObjectToObjectParam, ps.Sequence3(openParenthesis, namedObject, closeParenthesis)),
		ps.Map(anonymousObjectToObjectParam, ps.Sequence3(openParenthesis, anonymousObject, closeParenthesis)),
	)(xs)
}

var objectParamValue = ps.Or(
	ps.Map(stringToObjectParam, string_),
	ps.Map(numberToObjectParam, number),
)

var anonymousObjectToObjectParam = tuple.Fn3(func(_ tokenizer.Token, obj anonymousObjectNode, _ tokenizer.Token) st.ObjectParam {
	return st.NewRuleResultAnonymousObject(obj.Head(), obj.ParamPart())
})

func namedObjectWithoutChildToObjectParam(objName tokenizer.Token) st.ObjectParam {
	return st.NewRuleResultNamedObject(objName.Value(), nil)
}

var namedObjectToObjectParam = tuple.Fn3(func(_ tokenizer.Token, obj namedObjectNode, _ tokenizer.Token) st.ObjectParam {
	return st.NewRuleResultNamedObject(obj.Name(), obj.Params())
})

func numberToObjectParam(x tokenizer.Token) st.ObjectParam {
	return st.NewNumber(x.Value())
}

func stringToObjectParam(x tokenizer.Token) st.ObjectParam {
	s := x.Value()
	return st.NewString(s[1 : len(s)-1])
}

func variableToObjectParam(x tokenizer.Token) st.ObjectParam {
	return st.NewVariable(x.Value())
}
