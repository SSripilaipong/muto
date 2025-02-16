package pattern

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func nestedObjectRuleParamPattern(xs []psBase.Character) []tuple.Of2[base.PatternParam, []psBase.Character] {
	anonymousHead := ps.Or(
		psBase.BooleanPatternParam,
		psBase.StringPatternParam,
		psBase.NumberPatternParam,
		psBase.TagPatternParam,
		psBase.InParentheses(nestedObjectRuleParamPattern),
	)

	castAnonymousHead := func(head base.PatternParam) base.PatternParam {
		return stPattern.NewAnonymousRule(head, stPattern.ParamsToFixedParamPart([]base.PatternParam{}))
	}

	castAnonymousObjectWithParamPart := tuple.Fn2(func(head base.PatternParam, paramPart stPattern.ParamPart) base.PatternParam {
		return stPattern.NewAnonymousRule(head, paramPart)
	})

	return ps.Or(
		ps.Map(castAnonymousObjectWithParamPart, psBase.SpaceSeparated2(anonymousHead, rulePatternParamPart())),
		ps.Map(castAnonymousHead, anonymousHead),
		ps.Map(stPattern.NamedRuleToParam, Pattern()),
		ps.Map(stPattern.VariableRulePatternToRulePatternParam, variableRulePattern()),
	)(xs)
}
