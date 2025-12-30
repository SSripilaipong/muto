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

func TestParamPart_fixedParam(t *testing.T) {
	rule := ParamPart()

	t.Run("should parse rune as param", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`'x'`))
		expectedResult := stPattern.PatternsToParamPart([]stBase.Pattern{
			syntaxtree.NewRune(`'x'`),
		})
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse conjunction with object param", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`X^(g Y)`))
		expectedResult := stPattern.PatternsToParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				syntaxtree.NewVariable("X"),
				stPattern.NewNonDeterminantObject(
					syntaxtree.NewLocalClass("g"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
				),
			),
		})
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse conjunction with variable param", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`X^Y`))
		expectedResult := stPattern.PatternsToParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				syntaxtree.NewVariable("X"),
				syntaxtree.NewVariable("Y"),
			),
		})
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse conjunction with number param", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`X^1`))
		expectedResult := stPattern.PatternsToParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				syntaxtree.NewVariable("X"),
				syntaxtree.NewNumber("1"),
			),
		})
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})

	t.Run("should parse multiple conjunctions", func(t *testing.T) {
		r := rule(base.StringToCharTokens(`X^(A)^(B)`))
		expectedResult := stPattern.PatternsToParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				stPattern.NewConjunction(
					syntaxtree.NewVariable("X"),
					stPattern.NewNonDeterminantObject(
						syntaxtree.NewVariable("A"),
						stPattern.PatternsToFixedParamPart([]stBase.Pattern{}),
					),
				),
				stPattern.NewNonDeterminantObject(
					syntaxtree.NewVariable("B"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{}),
				),
			),
		})
		assert.Equal(t, expectedResult, ps.ResultValue(r))
		assert.Empty(t, r.X2())
	})
}
