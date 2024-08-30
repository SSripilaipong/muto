package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/common/tuple"
	"phi-lang/parser/tokenizer"
	"phi-lang/syntaxtree"
)

func object(xs []tokenizer.Token) []tuple.Of2[tuple.Of2[tokenizer.Token, []syntaxtree.ObjectParam], []tokenizer.Token] {
	return ps.Sequence2(objectName, ps.OptionalGreedyRepeat(objectParam))(xs)
}

func objectName(xs []tokenizer.Token) []tuple.Of2[tokenizer.Token, []tokenizer.Token] {
	return ps.Or(
		nonCapitalIdentifier,
		symbol,
	)(xs)
}

func objectParam(xs []tokenizer.Token) []tuple.Of2[syntaxtree.ObjectParam, []tokenizer.Token] {
	return ps.Or(
		objectParamValue,
		ps.Map(variableToObjectParam, variable),
		ps.Map(objectToObjectParam, ps.Sequence3(openParenthesis, object, closeParenthesis)),
	)(xs)
}

var objectParamValue = ps.Or(
	ps.Map(stringToObjectParam, string_),
	ps.Map(numberToObjectParam, number),
)

var objectToRuleResult = tuple.Fn2(func(name tokenizer.Token, params []syntaxtree.ObjectParam) syntaxtree.RuleResult {
	return syntaxtree.NewRuleResultObject(name.Value(), params)
})

var objectToObjectParam = tuple.Fn3(func(_ tokenizer.Token, obj tuple.Of2[tokenizer.Token, []syntaxtree.ObjectParam], _ tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewRuleResultObject(obj.X1().Value(), obj.X2())
})

func numberToObjectParam(x tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewNumber(x.Value())
}

func stringToObjectParam(x tokenizer.Token) syntaxtree.ObjectParam {
	s := x.Value()
	return syntaxtree.NewString(s[1 : len(s)-1])
}

func variableToObjectParam(x tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewVariable(x.Value())
}
