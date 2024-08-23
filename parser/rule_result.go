package parser

import (
	ps "phi-lang/common/parsing"
	st "phi-lang/parser/syntaxtree"
	"phi-lang/tokenizer"
)

var ruleResult = ps.Or(
	valueRuleResult,
	ps.Map(variableToRuleResult, variable),
	ps.Map(objectToRuleResult, object),
)

var valueRuleResult = ps.Or(
	ps.Map(stringToRuleResult, string_),
	ps.Map(numberToRuleResult, number),
)

func numberToRuleResult(x tokenizer.Token) st.RuleResult {
	return st.NewNumber(x.Value())
}

func stringToRuleResult(x tokenizer.Token) st.RuleResult {
	s := x.Value()
	return st.NewString(s[1 : len(s)-1])
}

func variableToRuleResult(x tokenizer.Token) st.RuleResult {
	return st.NewVariable(x.Value())
}
