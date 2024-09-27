package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func namedRulePattern() func(xs []tokenizer.Token) []tuple.Of2[stPattern.NamedRule, []tokenizer.Token] {
	cast := tuple.Fn2(func(name tokenizer.Token, params stPattern.ParamPart) stPattern.NamedRule {
		return stPattern.NewNamedRule(name.Value(), params)
	})

	return ps.Map(cast, ps.Sequence2(psBase.ClassIncludingSymbols, rulePatternParamPart()))
}

func variableRulePattern() func(xs []tokenizer.Token) []tuple.Of2[stPattern.VariableRule, []tokenizer.Token] {
	cast := tuple.Fn2(func(name tokenizer.Token, params stPattern.ParamPart) stPattern.VariableRule {
		return stPattern.NewVariableRulePattern(name.Value(), params)
	})

	return ps.Map(cast, ps.Sequence2(psBase.Variable, rulePatternParamPart()))
}

func rulePatternParamPart() func([]tokenizer.Token) []tuple.Of2[stPattern.ParamPart, []tokenizer.Token] {

	castLeftVariadic := tuple.Fn2(func(v variadicVarNode, p []stPattern.Param) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), p)
	})
	castRightVariadic := tuple.Fn2(func(p []stPattern.Param, v variadicVarNode) stPattern.ParamPart {
		return stPattern.NewRightVariadicParamPart(v.Name(), p)
	})

	fixedParam := fixedRuleParamPattern()
	return ps.Or(
		ps.Map(stPattern.ParamsToParamPart, ps.OptionalGreedyRepeat(fixedParam)),
		ps.Map(castLeftVariadic, ps.Sequence2(variadicVar, ps.OptionalGreedyRepeat(fixedParam))),
		ps.Map(castRightVariadic, ps.Sequence2(ps.GreedyRepeatAtLeastOnce(fixedParam), variadicVar)),
	)
}

func fixedRuleParamPattern() func(xs []tokenizer.Token) []tuple.Of2[stPattern.Param, []tokenizer.Token] {
	return ps.Or(
		ps.Map(variableToRuleParamPattern, psBase.Variable),
		ps.Map(booleanToRuleParamPattern, psBase.Boolean),
		ps.Map(stringToRuleParamPattern, psBase.String),
		ps.Map(numberToRuleParamPattern, psBase.Number),
		ps.Map(objectNameToRuleParamPattern, psBase.ClassIncludingSymbols),
		psBase.InParentheses(nestedObjectRuleParamPattern),
	)
}

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
