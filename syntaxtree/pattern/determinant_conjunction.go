package pattern

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type DeterminantConjunction struct {
	main  base.Determinant
	conjs []base.Pattern
}

func (DeterminantConjunction) PatternType() base.PatternType {
	return base.PatternTypeObject
}

func (DeterminantConjunction) DeterminantType() base.DeterminantType {
	return base.DeterminantTypeConjunction
}

func (c DeterminantConjunction) DeterminantName() string {
	return c.main.DeterminantName()
}

func (c DeterminantConjunction) Main() base.Determinant {
	return c.main
}

func (c DeterminantConjunction) Conjs() []base.Pattern {
	return c.conjs
}

func (c DeterminantConjunction) Head() base.Pattern {
	if base.IsPatternTypeObject(c.main) {
		return UnsafePatternToObject(c.main).Head()
	}
	return c.main
}

func (c DeterminantConjunction) ParamPart() ParamPart {
	if base.IsPatternTypeObject(c.main) {
		return UnsafePatternToObject(c.main).ParamPart()
	}
	return PatternsToFixedParamPart([]base.Pattern{})
}

func NewDeterminantConjunction(main base.Determinant, conjs []base.Pattern) DeterminantConjunction {
	return DeterminantConjunction{main: main, conjs: conjs}
}

// placeConjsInResult places the conjunctions at the correct level in the result slice.
// The level is determined by how many DeterminantObject levels are in the main.
func (c DeterminantConjunction) placeConjsInResult(result [][]base.Pattern) {
	level := countDeterminantObjectLayers(c.main)
	result[level] = append(result[level], c.conjs...)
}

func UnsafeDeterminantToConjunction(d base.Determinant) DeterminantConjunction {
	return d.(DeterminantConjunction)
}

func IsDeterminantConjunction(d base.Determinant) bool {
	return d.DeterminantType() == base.DeterminantTypeConjunction
}

func IsPatternDeterminantConjunction(p base.Pattern) bool {
	_, ok := p.(DeterminantConjunction)
	return ok
}

func UnsafePatternToDeterminantConjunction(p base.Pattern) DeterminantConjunction {
	return p.(DeterminantConjunction)
}

var _ base.Determinant = DeterminantConjunction{}
var _ Object = DeterminantConjunction{}
