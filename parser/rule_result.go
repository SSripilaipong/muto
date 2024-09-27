package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	"github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var ruleResult = ps.Or(
	valueRuleResult,
	ps.Map(variableToRuleResult, variable),
	ps.Map(namedObjectNodeToRuleResult, namedObject),
	ps.Map(anonymousObjectNodeToRuleResult, anonymousObject),
)

var valueRuleResult = ps.Or(
	ps.Map(booleanToRuleResult, boolean),
	ps.Map(stringToRuleResult, string_),
	ps.Map(numberToRuleResult, number),
)

func numberToRuleResult(x tokenizer.Token) stResult.Node {
	return syntaxtree.NewNumber(x.Value())
}

func booleanToRuleResult(x tokenizer.Token) stResult.Node {
	return syntaxtree.NewBoolean(x.Value())
}

func stringToRuleResult(x tokenizer.Token) stResult.Node {
	s := x.Value()
	return syntaxtree.NewString(s[1 : len(s)-1])
}

func variableToRuleResult(x tokenizer.Token) stResult.Node {
	return syntaxtree.NewVariable(x.Value())
}

var namedObjectNodeToRuleResult = func(obj namedObjectNode) stResult.Node {
	return stResult.NewNamedObject(obj.Name(), obj.Params())
}

var anonymousObjectNodeToRuleResult = func(obj anonymousObjectNode) stResult.Node {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}
