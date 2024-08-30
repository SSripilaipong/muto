package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/common/tuple"
	"phi-lang/parser/tokenizer"
	st "phi-lang/syntaxtree"
)

var rulePattern = ps.Map(mergeRulePattern, ps.Sequence2(identifier, ps.OptionalGreedyRepeat(ruleParamPattern)))

var ruleParamPattern = ps.Or(
	ps.Map(variableToRuleParamPattern, variable),
	ps.Map(stringToRuleParamPattern, string_),
	ps.Map(numberToRuleParamPattern, number),
)

var mergeRulePattern = tuple.Fn2(func(name tokenizer.Token, params []st.RuleParamPattern) st.RulePattern {
	return st.NewRulePattern(name.Value(), params)
})

func variableToRuleParamPattern(x tokenizer.Token) st.RuleParamPattern {
	return st.NewVariable(x.Value())
}

func stringToRuleParamPattern(x tokenizer.Token) st.RuleParamPattern {
	return st.NewString(x.Value())
}

func numberToRuleParamPattern(x tokenizer.Token) st.RuleParamPattern {
	return st.NewNumber(x.Value())
}
