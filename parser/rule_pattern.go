package parser

import (
	ps "phi-lang/common/parsing"
	"phi-lang/common/tuple"
	st "phi-lang/parser/syntaxtree"
	"phi-lang/tokenizer"
)

var rulePattern = ps.Map(mergeRulePattern, ps.Sequence2(identifier, ps.OptionalGreedyRepeat(ruleParamPattern)))

var ruleParamPattern = ps.Or(
	ps.Map(variableToRuleParamPattern, variable),
)

var mergeRulePattern = tuple.Fn2(func(name tokenizer.Token, params []st.RuleParamPattern) st.RulePattern {
	return st.NewRulePattern(name.Value(), params)
})

func variableToRuleParamPattern(x tokenizer.Token) st.RuleParamPattern {
	return st.NewVariable(x.Value())
}
