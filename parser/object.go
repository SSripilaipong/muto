package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/common/tuple"
	st "phi-lang/parser/syntaxtree"
	"phi-lang/parser/tokenizer"
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

var objectToRuleResult = tuple.Fn2(func(name tokenizer.Token, params []st.ObjectParam) st.RuleResult {
	return st.NewRuleResultObject(name.Value(), params)
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
