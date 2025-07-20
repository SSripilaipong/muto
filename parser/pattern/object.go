package pattern

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func Object() func(xs []psBase.Character) []tuple.Of2[stPattern.NonDeterminantObject, []psBase.Character] {

	castHead := func(head base.Pattern) stPattern.NonDeterminantObject {
		return stPattern.NewNonDeterminantObject(head, stPattern.PatternsToFixedParamPart([]base.Pattern{}))
	}
	castObject := tuple.Fn2(func(head base.Pattern, paramPart stPattern.ParamPart) stPattern.NonDeterminantObject {
		return stPattern.NewNonDeterminantObject(head, paramPart)
	})

	return func(xs []psBase.Character) []tuple.Of2[stPattern.NonDeterminantObject, []psBase.Character] {
		head := ps.First(
			psBase.FixedVarWithUnderscorePattern,
			psBase.BooleanPattern,
			psBase.StringPattern,
			psBase.RunePattern,
			psBase.NumberPattern,
			psBase.TagPattern,
			psBase.NonDeterminantClassRulePattern,
			ps.Map(base.ToPattern, Object()),
		)

		return psBase.InParentheses(ps.First(
			ps.Map(castObject, psBase.SpaceSeparated2(head, ParamPart())),
			ps.Map(castHead, head),
		))(xs)
	}
}
