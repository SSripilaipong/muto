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

	var parser func([]psBase.Character) tuple.Of2[rslt.Of[stPattern.DeterminantObject], []psBase.Character]
	parser = func(xs []psBase.Character) tuple.Of2[rslt.Of[stPattern.DeterminantObject], []psBase.Character] {
		headParser := ps.First(
			psBase.DeterminantClass,
			ps.Map(base.ToDeterminant, psBase.InParentheses(parser)),
		)
		withParam := ps.Map(castObject, psBase.SpaceSeparated2(headParser, ParamPart()))
		withoutParam := ps.Map(castHead, headParser)

		return ps.First(withParam, withoutParam)(xs)
	}
	return parser
}
