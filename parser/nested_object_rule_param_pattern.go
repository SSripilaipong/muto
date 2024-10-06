package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	tk "github.com/SSripilaipong/muto/parser/tokens"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func nestedObjectRuleParamPattern(xs []tk.Token) []tuple.Of2[stPattern.Param, []tk.Token] {
	anonymousHead := ps.Or(
		ps.Map(booleanToRuleParamPattern, psBase.Boolean),
		ps.Map(stringToRuleParamPattern, psBase.String),
		ps.Map(numberToRuleParamPattern, psBase.Number),
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
