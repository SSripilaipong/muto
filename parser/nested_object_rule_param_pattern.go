package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func nestedObjectRuleParamPattern(xs []psBase.Character) []tuple.Of2[stPattern.Param, []psBase.Character] {
	anonymousHead := ps.Or(
		psBase.BooleanPatternParam,
		psBase.StringPatternParam,
		psBase.NumberPatternParam,
		psBase.InParentheses(nestedObjectRuleParamPattern),
	)

	castAnonymousHead := func(head stPattern.Param) stPattern.Param {
		return stPattern.NewAnonymousRule(head, stPattern.ParamsToFixedParamPart([]stPattern.Param{}))
	}

	castAnonymousObjectWithParamPart := tuple.Fn2(func(head stPattern.Param, paramPart stPattern.ParamPart) stPattern.Param {
		return stPattern.NewAnonymousRule(head, paramPart)
	})

	return ps.Or(
		ps.Map(castAnonymousObjectWithParamPart, psBase.SpaceSeparated2(anonymousHead, rulePatternParamPart())),
		ps.Map(castAnonymousHead, anonymousHead),
		ps.Map(stPattern.NamedRuleToParam, namedRulePattern()),
		ps.Map(stPattern.VariableRulePatternToRulePatternParam, variableRulePattern()),
	)(xs)
}
