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
			ps.ToParser(psBase.DeterminantClass),
			ps.Map(base.ToDeterminant, ps.ToParser(psBase.InParentheses(parser))),
		).Legacy
		withParam := ps.Map(castObject, ps.ToParser(psBase.SpaceSeparated2(headParser, ParamPart())))
		withoutParam := ps.Map(castHead, ps.ToParser(headParser))

		return ps.First(withParam, withoutParam).Legacy(xs)
	}
	return parser
}
