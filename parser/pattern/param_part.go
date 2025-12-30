package pattern

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func ParamPart() (parser func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.ParamPart], []psBase.Character]) {
	corePattern := buildCorePatternParser()
	patternWithConjunction := buildConjunctionPatternParser(corePattern)
	parser = func(xs []psBase.Character) tuple.Of2[rslt.Of[stPattern.ParamPart], []psBase.Character] {
		return paramPart(patternWithConjunction)(xs)
	}
	return
}

func ParamPartWithoutConjunction() func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.ParamPart], []psBase.Character] {
	corePattern := buildCorePatternParser()
	return paramPart(corePattern) // should this really exist?
}

func buildCorePatternParser() func([]psBase.Character) tuple.Of2[rslt.Of[base.Pattern], []psBase.Character] {
	objectContent := ObjectContent()
	return fixedParam(ObjectParenthesis(objectContent))
}

func paramPart(
	fixedParamParser func([]psBase.Character) tuple.Of2[rslt.Of[base.Pattern], []psBase.Character],
) func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.ParamPart], []psBase.Character] {
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

func fixedParam(
	object func(xs []psBase.Character) tuple.Of2[rslt.Of[stPattern.NonDeterminantObject], []psBase.Character],
) func([]psBase.Character) tuple.Of2[rslt.Of[base.Pattern], []psBase.Character] {
	return ps.First(
		ps.ToParser(psBase.FixedVarWithUnderscorePattern),
		ps.ToParser(psBase.BooleanPattern),
		ps.ToParser(psBase.StringPattern),
		ps.ToParser(psBase.RunePattern),
		ps.ToParser(psBase.NumberPattern),
		ps.ToParser(psBase.TagPattern),
		ps.ToParser(psBase.NonDeterminantClassRulePattern),
		ps.Map(base.ToPattern, ps.ToParser(object)),
	).Legacy
}

func buildConjunctionPatternParser(
	corePattern func([]psBase.Character) tuple.Of2[rslt.Of[base.Pattern], []psBase.Character],
) func([]psBase.Character) tuple.Of2[rslt.Of[base.Pattern], []psBase.Character] {
	mergeConjunctions := tuple.Fn2(func(main base.Pattern, conjs []base.Pattern) base.Pattern {
		if len(conjs) == 0 { // not needed?
			return main
		}
		result := main
		for _, conj := range conjs {
			result = stPattern.NewConjunction(result, conj)
		}
		return result
	})
	caretPattern := ps.Map(tuple.Fn2(func(_ string, p base.Pattern) base.Pattern { return p }),
		ps.ToParser(psBase.IgnoreWhiteSpaceBetween2(
			psBase.FixedChars("^"),
			corePattern,
		)),
	)

	return ps.Map(mergeConjunctions,
		ps.Sequence2(ps.ToParser(corePattern), ps.OptionalGreedyRepeat(caretPattern)),
	).Legacy
}
