package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	st "muto/syntaxtree"
)

func rulePattern() func(xs []tokenizer.Token) []tuple.Of2[st.RulePattern, []tokenizer.Token] {
	return ps.Map(mergeRulePattern, ps.Sequence2(identifier, ps.OptionalGreedyRepeat(ruleParamPattern())))
}

func ruleParamPattern() func(xs []tokenizer.Token) []tuple.Of2[st.RuleParamPattern, []tokenizer.Token] {
	return ps.Or(
		ps.Map(variableToRuleParamPattern, variable),
		ps.Map(stringToRuleParamPattern, string_),
		ps.Map(numberToRuleParamPattern, number),
		nestedObjectRuleParamPattern,
	)
}

func mergeRulePattern(xs tuple.Of2[tokenizer.Token, []st.RuleParamPattern]) st.RulePattern {
	return tuple.Fn2(func(name tokenizer.Token, params []st.RuleParamPattern) st.RulePattern {
		return st.NewRulePattern(name.Value(), params)
	})(xs)
}

func nestedObjectRuleParamPattern(xs []tokenizer.Token) []tuple.Of2[st.RuleParamPattern, []tokenizer.Token] {
	pattern := ps.Sequence3(openParenthesis, rulePattern(), closeParenthesis)
	return ps.Map(mergeNestedObjectRuleParamPattern, pattern)(xs)
}

var mergeNestedObjectRuleParamPattern = tuple.Fn3(func(_ tokenizer.Token, o st.RulePattern, _ tokenizer.Token) st.RuleParamPattern {
	return o
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
