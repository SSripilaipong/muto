package parser

import (
	ps "muto/common/parsing"
	"muto/parser/tokenizer"
	"muto/syntaxtree"
)

var ruleResult = ps.Or(
	valueRuleResult,
	ps.Map(variableToRuleResult, variable),
	ps.Map(namedObjectNodeToRuleResult, namedObject),
	ps.Map(anonymousObjectNodeToRuleResult, anonymousObject),
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

var namedObjectNodeToRuleResult = func(obj namedObjectNode) syntaxtree.RuleResult {
	return syntaxtree.NewRuleResultNamedObject(obj.Name(), obj.Params())
}

var anonymousObjectNodeToRuleResult = func(obj anonymousObjectNode) syntaxtree.RuleResult {
	return syntaxtree.NewRuleResultAnonymousObject(obj.Head(), obj.ParamPart())
}
