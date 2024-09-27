package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var ruleResult = ps.Or(
	valueRuleResult,
	ps.Map(variableToRuleResult, psBase.Variable),
	ps.Map(objectNodeToRuleResult, object),
)

var valueRuleResult = ps.Or(
	ps.Map(booleanToRuleResult, psBase.Boolean),
	ps.Map(stringToRuleResult, psBase.String),
	ps.Map(numberToRuleResult, psBase.Number),
	ps.Map(classToRuleResult, psBase.ClassIncludingSymbols),
)

func classToRuleResult(x tokenizer.Token) stResult.Node {
	return st.NewClass(x.Value())
}

func numberToRuleResult(x tokenizer.Token) stResult.Node {
	return st.NewNumber(x.Value())
}

func booleanToRuleResult(x tokenizer.Token) stResult.Node {
	return st.NewBoolean(x.Value())
}

func stringToRuleResult(x tokenizer.Token) stResult.Node {
	s := x.Value()
	return st.NewString(s[1 : len(s)-1])
}

func variableToRuleResult(x tokenizer.Token) stResult.Node {
	return st.NewVariable(x.Value())
}

var objectNodeToRuleResult = func(obj objectNode) stResult.Node {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}
