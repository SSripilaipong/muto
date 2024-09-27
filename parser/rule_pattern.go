package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func namedRulePattern() func(xs []tokenizer.Token) []tuple.Of2[stPattern.NamedRule, []tokenizer.Token] {
	return ps.Map(mergeNamedRulePattern, ps.Sequence2(classIncludingSymbols, rulePatternParamPart()))
}

func variableRulePattern() func(xs []tokenizer.Token) []tuple.Of2[stPattern.VariableRule, []tokenizer.Token] {
	return ps.Map(mergeVariableRulePattern, ps.Sequence2(variable, rulePatternParamPart()))
}

func rulePatternParamPart() func([]tokenizer.Token) []tuple.Of2[stPattern.ParamPart, []tokenizer.Token] {
	fixedParam := fixedRuleParamPattern()
	return ps.Or(
		ps.Map(stPattern.ParamsToParamPart, ps.OptionalGreedyRepeat(fixedParam)),
		ps.Map(mergeLeftVariadicRulePatternParamPart, ps.Sequence2(variadicVar, ps.OptionalGreedyRepeat(fixedParam))),
		ps.Map(mergeRightVariadicRulePatternParamPart, ps.Sequence2(ps.GreedyRepeatAtLeastOnce(fixedParam), variadicVar)),
	)
}

var mergeLeftVariadicRulePatternParamPart = tuple.Fn2(func(v variadicVarNode, p []stPattern.Param) stPattern.ParamPart {
	return stPattern.NewLeftVariadicParamPart(v.Name(), p)
})

var mergeRightVariadicRulePatternParamPart = tuple.Fn2(func(p []stPattern.Param, v variadicVarNode) stPattern.ParamPart {
	return stPattern.NewRightVariadicParamPart(v.Name(), p)
})

func fixedRuleParamPattern() func(xs []tokenizer.Token) []tuple.Of2[stPattern.Param, []tokenizer.Token] {
	return ps.Or(
		ps.Map(variableToRuleParamPattern, variable),
		ps.Map(booleanToRuleParamPattern, boolean),
		ps.Map(stringToRuleParamPattern, string_),
		ps.Map(numberToRuleParamPattern, number),
		ps.Map(objectNameToRuleParamPattern, classIncludingSymbols),
		nestedObjectRuleParamPattern,
	)
}

var mergeNamedRulePattern = tuple.Fn2(func(name tokenizer.Token, params stPattern.ParamPart) stPattern.NamedRule {
	return stPattern.NewNamedRule(name.Value(), params)
})

var mergeVariableRulePattern = tuple.Fn2(func(name tokenizer.Token, params stPattern.ParamPart) stPattern.VariableRule {
	return stPattern.NewVariableRulePattern(name.Value(), params)
})

func variableToRuleParamPattern(x tokenizer.Token) stPattern.Param {
	return st.NewVariable(x.Value())
}

func booleanToRuleParamPattern(x tokenizer.Token) stPattern.Param {
	return st.NewBoolean(x.Value())
}

func stringToRuleParamPattern(x tokenizer.Token) stPattern.Param {
	return st.NewString(x.Value())
}

func numberToRuleParamPattern(x tokenizer.Token) stPattern.Param {
	return st.NewNumber(x.Value())
}

func objectNameToRuleParamPattern(x tokenizer.Token) stPattern.Param {
	return stPattern.NewNamedRule(x.Value(), stPattern.FixedParamPart{})
}
