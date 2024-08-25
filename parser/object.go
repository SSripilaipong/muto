package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/common/tuple"
	"phi-lang/parser/tokenizer"
	"phi-lang/syntaxtree"
)

var object = ps.Sequence2(objectName, ps.OptionalGreedyRepeat(objectParam))

var objectName = ps.Or(
	nonCapitalIdentifier,
	symbol,
)

var objectParam = ps.Or(
	objectParamValue,
	ps.Map(variableToObjectParam, variable),
)

var objectParamValue = ps.Or(
	ps.Map(stringToObjectParam, string_),
	ps.Map(numberToObjectParam, number),
)

var objectToRuleResult = tuple.Fn2(func(name tokenizer.Token, params []syntaxtree.ObjectParam) syntaxtree.RuleResult {
	return syntaxtree.NewRuleResultObject(name.Value(), params)
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
