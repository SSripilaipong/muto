package pattern

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func ParamPart() func([]psBase.Character) []tuple.Of2[stPattern.ParamPart, []psBase.Character] {

	fixedParam := ps.Or(
		psBase.FixedVarWithUnderscorePattern,
		psBase.BooleanPattern,
		psBase.StringPattern,
		psBase.NumberPattern,
		psBase.TagPattern,
		psBase.ClassRulePattern,
		ps.Map(base.ToPattern, Object()),
	)

	castVariadic := func(v psBase.VariadicVarNode) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), stPattern.PatternsToFixedParamPart([]base.Pattern{}))
	}
	castLeftVariadic := tuple.Fn2(func(v psBase.VariadicVarNode, p []base.Pattern) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), p)
	})
	castRightVariadic := tuple.Fn2(func(p []base.Pattern, v psBase.VariadicVarNode) stPattern.ParamPart {
		return stPattern.NewRightVariadicParamPart(v.Name(), p)
	})

	return ps.Or(
		ps.Map(castVariadic, psBase.VariadicVarWithUnderscore),
		ps.Map(castLeftVariadic, psBase.SpaceSeparated2(psBase.VariadicVarWithUnderscore, psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParam))),
		ps.Map(castRightVariadic, psBase.SpaceSeparated2(psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParam), psBase.VariadicVarWithUnderscore)),
		ps.Map(stPattern.PatternsToParamPart, psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParam)),
	)
}
