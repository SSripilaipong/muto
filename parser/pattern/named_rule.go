package pattern

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func variableRulePattern() func(xs []psBase.Character) []tuple.Of2[stPattern.VariableRule, []psBase.Character] {
	cast := tuple.Fn2(func(name base.Variable, params stPattern.ParamPart) stPattern.VariableRule {
		return stPattern.NewVariableRulePattern(name.Name(), params)
	})

	return ps.Map(cast, psBase.SpaceSeparated2(psBase.FixedVar, rulePatternParamPart()))
}

func rulePatternParamPart() func([]psBase.Character) []tuple.Of2[stPattern.ParamPart, []psBase.Character] {

	castVariadic := func(v psBase.VariadicVarNode) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), stPattern.ParamsToFixedParamPart([]base.PatternParam{}))
	}
	castLeftVariadic := tuple.Fn2(func(v psBase.VariadicVarNode, p []base.PatternParam) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), p)
	})
	castRightVariadic := tuple.Fn2(func(p []base.PatternParam, v psBase.VariadicVarNode) stPattern.ParamPart {
		return stPattern.NewRightVariadicParamPart(v.Name(), p)
	})

	fixedParam := fixedRuleParamPattern()
	return ps.Or(
		ps.Map(stPattern.ParamsToParamPart, psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParam)),
		ps.Map(castVariadic, psBase.VariadicVar),
		ps.Map(castLeftVariadic, psBase.SpaceSeparated2(psBase.VariadicVar, psBase.OptionalGreedyRepeatSpaceSeparated(fixedParam))),
		ps.Map(castRightVariadic, psBase.SpaceSeparated2(psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParam), psBase.VariadicVar)),
	)
}

func fixedRuleParamPattern() func(xs []psBase.Character) []tuple.Of2[base.PatternParam, []psBase.Character] {
	var classToPatternParam = func(x base.Class) base.PatternParam {
		return stPattern.NewNamedRule(x.Value(), stPattern.FixedParamPart{})
	}

	return ps.Or(
		psBase.FixedVarPatternParam,
		psBase.BooleanPatternParam,
		psBase.StringPatternParam,
		psBase.NumberPatternParam,
		psBase.TagPatternParam,
		ps.Map(classToPatternParam, psBase.ClassRule),
		psBase.InParentheses(nestedObjectRuleParamPattern()),
	)
}
