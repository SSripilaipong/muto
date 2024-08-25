package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/parser/tokenizer"
	"phi-lang/syntaxtree"
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

func numberToRuleResult(x tokenizer.Token) syntaxtree.RuleResult {
	return syntaxtree.NewNumber(x.Value())
}

func stringToRuleResult(x tokenizer.Token) syntaxtree.RuleResult {
	s := x.Value()
	return syntaxtree.NewString(s[1 : len(s)-1])
}

func variableToRuleResult(x tokenizer.Token) syntaxtree.RuleResult {
	return syntaxtree.NewVariable(x.Value())
}
