package pattern

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func Determinant() func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.DeterminantObject], []psBase.Character] {

	castHead := func(head base.Determinant) stPattern.DeterminantObject {
		return stPattern.NewDeterminantObject(head, stPattern.PatternsToFixedParamPart([]base.Pattern{}))
	}
	castObject := tuple.Fn2(func(head base.Determinant, paramPart stPattern.ParamPart) stPattern.DeterminantObject {
		return stPattern.NewDeterminantObject(head, paramPart)
	})

	var self func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.DeterminantObject], []psBase.Character]
	self = func(xs []psBase.Character) tuple.Of2[rslt.Of[stPattern.DeterminantObject], []psBase.Character] {
		// Parse head with optional conjunctions
		headWithConj := determinantConjunction(self)

		withParam := ps.Map(castObject, ps.ToParser(psBase.SpaceSeparated2(headWithConj, ParamPart())))
		withoutParam := ps.Map(castHead, ps.ToParser(headWithConj))

		return ps.First(withParam, withoutParam).Legacy(xs)
	}
	return self
}

// determinantConjunction wraps a determinant parser to support conjunction syntax.
// It parses a determinant head followed by optional ^pattern sequences.
func determinantConjunction(
	determinant func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.DeterminantObject], []psBase.Character],
) func([]psBase.Character) tuple.Of2[rslt.Of[base.Determinant], []psBase.Character] {

	// Parser for ^pattern
	caretPattern := ps.Map(
		tuple.Fn2(func(_ psBase.Character, p base.Pattern) base.Pattern { return p }),
		ps.ToParser(psBase.IgnoreWhiteSpaceBetween2(
			psBase.Caret, buildCorePatternParser(),
		)),
	)

	// Wrap determinant with conjunctions if present
	castConjunction := tuple.Fn2(func(det base.Determinant, conjs []base.Pattern) base.Determinant {
		if len(conjs) == 0 {
			return det
		}
		return stPattern.NewDeterminantConjunction(det, conjs)
	})

	// Base head parser: either a class or a parenthesized determinant
	head := ps.First(
		ps.ToParser(psBase.DeterminantClass),
		ps.Map(base.ToDeterminant, ps.ToParser(psBase.InParentheses(determinant))),
	).Legacy

	// Head with optional conjunctions
	return ps.Map(
		castConjunction,
		ps.Sequence2(ps.ToParser(head), ps.OptionalGreedyRepeat(caretPattern)),
	).Legacy
}
