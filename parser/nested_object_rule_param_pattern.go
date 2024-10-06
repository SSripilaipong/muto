package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func nestedObjectRuleParamPattern(xs []tk.Token) []tuple.Of2[stPattern.Param, []tk.Token] {
	anonymousHead := ps.Or(
		ps.Map(booleanToRuleParamPattern, psBase.Boolean),
		ps.Map(stringToRuleParamPattern, psBase.String),
		ps.Map(numberToRuleParamPattern, psBase.Number),
		psBase.InParentheses(nestedObjectRuleParamPattern),
	)

	castAnonymousObject := tuple.Fn2(func(head stPattern.Param, paramPart stPattern.ParamPart) stPattern.Param {
		return stPattern.NewAnonymousRule(head, paramPart)
	})

	return ps.Or(
		ps.Map(castAnonymousObject, psBase.SpaceSeparated2(anonymousHead, rulePatternParamPart())),
		ps.Map(stPattern.NamedRuleToParam, namedRulePattern()),
		ps.Map(stPattern.VariableRulePatternToRulePatternParam, variableRulePattern()),
	)(xs)
}
