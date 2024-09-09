package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	st "muto/syntaxtree"
)

func namedRulePattern() func(xs []tokenizer.Token) []tuple.Of2[st.NamedRulePattern, []tokenizer.Token] {
	return ps.Map(mergeNamedRulePattern, ps.Sequence2(objectName, rulePatternParamPart()))
}

func variableRulePattern() func(xs []tokenizer.Token) []tuple.Of2[st.VariableRulePattern, []tokenizer.Token] {
	return ps.Map(mergeVariableRulePattern, ps.Sequence2(variable, rulePatternParamPart()))
}

func rulePatternParamPart() func([]tokenizer.Token) []tuple.Of2[st.RulePatternParamPart, []tokenizer.Token] {
	fixedParam := fixedRuleParamPattern()
	return ps.Or(
		ps.Map(st.RuleParamPatternsToRulePatternParamPart, ps.OptionalGreedyRepeat(fixedParam)),
		ps.Map(mergeLeftVariadicRulePatternParamPart, ps.Sequence2(variadicVar, ps.OptionalGreedyRepeat(fixedParam))),
		ps.Map(mergeRightVariadicRulePatternParamPart, ps.Sequence2(ps.GreedyRepeatAtLeastOnce(fixedParam), variadicVar)),
	)
}

var mergeLeftVariadicRulePatternParamPart = tuple.Fn2(func(v variadicVarNode, p []st.RuleParamPattern) st.RulePatternParamPart {
	return st.NewRulePatternLeftVariadicParamPart(v.Name(), p)
})

var mergeRightVariadicRulePatternParamPart = tuple.Fn2(func(p []st.RuleParamPattern, v variadicVarNode) st.RulePatternParamPart {
	return st.NewRulePatternRightVariadicParamPart(v.Name(), p)
})

func fixedRuleParamPattern() func(xs []tokenizer.Token) []tuple.Of2[st.RuleParamPattern, []tokenizer.Token] {
	return ps.Or(
		ps.Map(variableToRuleParamPattern, variable),
		ps.Map(stringToRuleParamPattern, string_),
		ps.Map(numberToRuleParamPattern, number),
		ps.Map(objectNameToRuleParamPattern, objectName),
		nestedObjectRuleParamPattern,
	)
}

var mergeNamedRulePattern = tuple.Fn2(func(name tokenizer.Token, params st.RulePatternParamPart) st.NamedRulePattern {
	return st.NewNamedRulePattern(name.Value(), params)
})

var mergeVariableRulePattern = tuple.Fn2(func(name tokenizer.Token, params st.RulePatternParamPart) st.VariableRulePattern {
	return st.NewVariableRulePattern(name.Value(), params)
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

func objectNameToRuleParamPattern(x tokenizer.Token) st.RuleParamPattern {
	return st.NewNamedRulePattern(x.Value(), st.RulePatternFixedParamPart{})
}
