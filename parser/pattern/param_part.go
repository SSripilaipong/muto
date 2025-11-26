package pattern

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func ParamPart() func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.ParamPart], []psBase.Character] {

	fixedParamParser := ps.First(
		ps.ToParser(psBase.FixedVarWithUnderscorePattern),
		ps.ToParser(psBase.BooleanPattern),
		ps.ToParser(psBase.StringPattern),
		ps.ToParser(psBase.RunePattern),
		ps.ToParser(psBase.NumberPattern),
		ps.ToParser(psBase.TagPattern),
		ps.ToParser(psBase.NonDeterminantClassRulePattern),
		ps.Map(base.ToPattern, ps.ToParser(Object())),
	).Legacy

	castVariadic := func(v psBase.VariadicVarNode) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), stPattern.PatternsToFixedParamPart([]base.Pattern{}))
	}
	castLeftVariadic := tuple.Fn2(func(v psBase.VariadicVarNode, p []base.Pattern) stPattern.ParamPart {
		return stPattern.NewLeftVariadicParamPart(v.Name(), p)
	})
	castRightVariadic := tuple.Fn2(func(p []base.Pattern, v psBase.VariadicVarNode) stPattern.ParamPart {
		return stPattern.NewRightVariadicParamPart(v.Name(), p)
	})

	return ps.First(
		ps.Map(
			castLeftVariadic,
			ps.ToParser(psBase.SpaceSeparated2(
				psBase.VariadicVarWithUnderscore,
				psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParamParser),
			)),
		),
		ps.Map(
			castVariadic,
			ps.ToParser(psBase.VariadicVarWithUnderscore),
		),
		ps.Map(
			castRightVariadic,
			ps.ToParser(psBase.SpaceSeparated2(
				psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParamParser),
				psBase.VariadicVarWithUnderscore,
			)),
		),
		ps.Map(
			stPattern.PatternsToParamPart,
			ps.ToParser(psBase.GreedyRepeatAtLeastOnceSpaceSeparated(fixedParamParser)),
		),
	).Legacy
}
