package pattern

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func Object() func(xs []psBase.Character) tuple.Of2[rslt.Of[stPattern.NonDeterminantObject], []psBase.Character] {
	castHead := func(head base.Pattern) stPattern.NonDeterminantObject {
		return stPattern.NewNonDeterminantObject(head, stPattern.PatternsToFixedParamPart([]base.Pattern{}))
	}
	castObject := tuple.Fn2(func(head base.Pattern, paramPart stPattern.ParamPart) stPattern.NonDeterminantObject {
		return stPattern.NewNonDeterminantObject(head, paramPart)
	})

	var parser func(xs []psBase.Character) tuple.Of2[rslt.Of[stPattern.NonDeterminantObject], []psBase.Character]
	parser = func(xs []psBase.Character) tuple.Of2[rslt.Of[stPattern.NonDeterminantObject], []psBase.Character] {
		headParser := ps.First(
			psBase.FixedVarWithUnderscorePattern,
			psBase.BooleanPattern,
			psBase.StringPattern,
			psBase.RunePattern,
			psBase.NumberPattern,
			psBase.TagPattern,
			psBase.NonDeterminantClassRulePattern,
			ps.Map(base.ToPattern, parser),
		)

		return psBase.InParentheses(ps.First(
			ps.Map(castObject, psBase.SpaceSeparated2(
				headParser,
				ParamPart(),
			)),
			ps.Map(castHead, headParser),
		))(xs)
	}
	return parser
}
