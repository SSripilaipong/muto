package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestDeterminant_tag(t *testing.T) {
	rule := Determinant()

	t.Run("should parse tag as a child", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`f 1 .abc 2`))
		expectedResult := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToParamPart([]stBase.Pattern{
			syntaxtree.NewNumber("1"), syntaxtree.NewTag(`.abc`), syntaxtree.NewNumber("2"),
		}))
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse tag as a child of a nested named rule", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`f (1 .abc 2)`))
		expectedResult := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(syntaxtree.NewNumber("1"), stPattern.PatternsToParamPart([]stBase.Pattern{
				syntaxtree.NewTag(`.abc`), syntaxtree.NewNumber("2"),
			})),
		}))
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})
}

func TestDeterminant_variadic(t *testing.T) {
	rule := Determinant()

	t.Run("should parse nested variadic param", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`g (f Xs...)`))
		expectedResult := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.PatternsToParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(syntaxtree.NewLocalClass("f"),
				stPattern.NewLeftVariadicParamPart("Xs", stPattern.FixedParamPart{})),
		}))
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})
}

func TestDeterminant_nested_object(t *testing.T) {
	rule := Determinant()

	t.Run("should match empty nested head", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`(f)`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToParamPart([]stBase.Pattern{})),
			stPattern.PatternsToParamPart([]stBase.Pattern{}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should match empty nested head with outer params", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`(f) 1`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToParamPart([]stBase.Pattern{})),
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewNumber("1")}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should match nested head with inner params", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`(f 1) 2`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantObject(
				syntaxtree.NewLocalClass("f"),
				stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewNumber("1")}),
			),
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewNumber("2")}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})
}

func TestDeterminant_conjunction(t *testing.T) {
	rule := Determinant()

	t.Run("should parse simple determinant conjunction", func(t *testing.T) {
		// Pattern: (f X)^P
		r := rule(base.StringToCharTokens(`(f X)^P`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
				),
				[]stBase.Pattern{syntaxtree.NewVariable("P")},
			),
			stPattern.PatternsToParamPart([]stBase.Pattern{}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse determinant conjunction with outer params", func(t *testing.T) {
		// Pattern: (f X)^P Y
		r := rule(base.StringToCharTokens(`(f X)^P Y`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
				),
				[]stBase.Pattern{syntaxtree.NewVariable("P")},
			),
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse multiple conjunctions at same level", func(t *testing.T) {
		// Pattern: (f X)^P^Q
		r := rule(base.StringToCharTokens(`(f X)^P^Q`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
				),
				[]stBase.Pattern{syntaxtree.NewVariable("P"), syntaxtree.NewVariable("Q")},
			),
			stPattern.PatternsToParamPart([]stBase.Pattern{}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse nested conjunctions at different levels", func(t *testing.T) {
		// Pattern: ((f X)^P Y)^Q Z
		r := rule(base.StringToCharTokens(`((f X)^P Y)^Q Z`))
		innerConj := stPattern.NewDeterminantConjunction(
			stPattern.NewDeterminantObject(
				syntaxtree.NewLocalClass("f"),
				stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
			),
			[]stBase.Pattern{syntaxtree.NewVariable("P")},
		)
		middleObj := stPattern.NewDeterminantObject(
			innerConj,
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
		)
		outerConj := stPattern.NewDeterminantConjunction(
			middleObj,
			[]stBase.Pattern{syntaxtree.NewVariable("Q")},
		)
		expectedResult := stPattern.NewDeterminantObject(
			outerConj,
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("Z")}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse bare class with conjunction", func(t *testing.T) {
		// Pattern: f^P X
		r := rule(base.StringToCharTokens(`f^P X`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				syntaxtree.NewLocalClass("f"),
				[]stBase.Pattern{syntaxtree.NewVariable("P")},
			),
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse conjunction with object pattern", func(t *testing.T) {
		// Pattern: (f X)^(g Y) Z
		r := rule(base.StringToCharTokens(`(f X)^(g Y) Z`))
		expectedResult := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
				),
				[]stBase.Pattern{
					stPattern.NewNonDeterminantObject(
						syntaxtree.NewLocalClass("g"),
						stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
					),
				},
			),
			stPattern.PatternsToParamPart([]stBase.Pattern{syntaxtree.NewVariable("Z")}),
		)
		assert.True(t, ps.IsResultOk(r))
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})
}
