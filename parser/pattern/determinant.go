package pattern

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

var RsDeterminant = ps.Map(rslt.Value, Determinant())

func Determinant() func(xs []psBase.Character) []tuple.Of2[stPattern.DeterminantObject, []psBase.Character] {

	castHead := func(head base.Determinant) stPattern.DeterminantObject {
		return stPattern.NewDeterminantObject(head, stPattern.PatternsToFixedParamPart([]base.Pattern{}))
	}
	castObject := tuple.Fn2(func(head base.Determinant, paramPart stPattern.ParamPart) stPattern.DeterminantObject {
		return stPattern.NewDeterminantObject(head, paramPart)
	})

	return func(xs []psBase.Character) []tuple.Of2[stPattern.DeterminantObject, []psBase.Character] {
		head := ps.First(
			psBase.ClassDeterminant,
			ps.Map(base.ToDeterminant, psBase.InParentheses(Determinant())),
		)

		return ps.First(
			ps.Map(castObject, psBase.SpaceSeparated2(head, ParamPart())),
			ps.Map(castHead, head),
		)(xs)
	}
}
