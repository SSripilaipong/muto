package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func namedRulePattern() func(xs []tk.Token) []tuple.Of2[stPattern.NamedRule, []tk.Token] {
	cast := tuple.Fn2(func(name tk.Token, params stPattern.ParamPart) stPattern.NamedRule {
		return stPattern.NewNamedRule(name.Value(), params)
	})

	return ps.Map(cast, psBase.SpaceSeparated2(psBase.Class, rulePatternParamPart()))
}

func variableRulePattern() func(xs []tk.Token) []tuple.Of2[stPattern.VariableRule, []tk.Token] {
	cast := tuple.Fn2(func(name tk.Token, params stPattern.ParamPart) stPattern.VariableRule {
		return stPattern.NewVariableRulePattern(name.Value(), params)
	})

	return ps.Map(cast, psBase.SpaceSeparated2(psBase.FixedVar, rulePatternParamPart()))
}

func rulePatternParamPart() func([]tk.Token) []tuple.Of2[stPattern.ParamPart, []tk.Token] {

	castLeftVariadic := tuple.Fn2(func(v psBase.VariadicVarNode, p []stPattern.Param) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), p)
	})
	castRightVariadic := tuple.Fn2(func(p []stPattern.Param, v psBase.VariadicVarNode) stPattern.ParamPart {
		return stPattern.NewRightVariadicParamPart(v.Name(), p)
	})

	fixedParam := fixedRuleParamPattern()
	return ps.Or(
		ps.Map(stPattern.ParamsToParamPart, ps.OptionalGreedyRepeat(fixedParam)),
		ps.Map(castLeftVariadic, psBase.SpaceSeparated2(psBase.VariadicVar, ps.OptionalGreedyRepeat(fixedParam))),
		ps.Map(castRightVariadic, psBase.SpaceSeparated2(ps.GreedyRepeatAtLeastOnce(fixedParam), psBase.VariadicVar)),
	)
}

func fixedRuleParamPattern() func(xs []tk.Token) []tuple.Of2[stPattern.Param, []tk.Token] {
	return ps.Or(
		ps.Map(variableToRuleParamPattern, psBase.FixedVar),
		ps.Map(booleanToRuleParamPattern, psBase.Boolean),
		ps.Map(stringToRuleParamPattern, psBase.String),
		ps.Map(numberToRuleParamPattern, psBase.Number),
		ps.Map(objectNameToRuleParamPattern, psBase.Class),
		psBase.InParentheses(nestedObjectRuleParamPattern),
	)
}

func variableToRuleParamPattern(x tk.Token) stPattern.Param {
	return st.NewVariable(x.Value())
}

func booleanToRuleParamPattern(x tk.Token) stPattern.Param {
	return st.NewBoolean(x.Value())
}

func stringToRuleParamPattern(x tk.Token) stPattern.Param {
	return st.NewString(x.Value())
}

func numberToRuleParamPattern(x tk.Token) stPattern.Param {
	return st.NewNumber(x.Value())
}

func objectNameToRuleParamPattern(x tk.Token) stPattern.Param {
	return stPattern.NewNamedRule(x.Value(), stPattern.FixedParamPart{})
}
