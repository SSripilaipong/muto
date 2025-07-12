package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariadicVar(t *testing.T) {
	t.Run("should parse single letter variadic variable", func(t *testing.T) {
		r := VariadicVarWithUnderscore(StringToCharTokens(`X...123`))
		expectedResult := newVariadicVar("X")
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse multiple letters variadic variable", func(t *testing.T) {
		r := VariadicVarWithUnderscore(StringToCharTokens(`XYabc-123!'?-1s...123`))
		expectedResult := newVariadicVar(`XYabc-123!'?-1s`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := VariadicVarWithUnderscore(StringToCharTokens(`xy...`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse fixed var", func(t *testing.T) {
		r := VariadicVarWithUnderscore(StringToCharTokens(`X.123`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse variadic var with . in the name", func(t *testing.T) {
		r := VariadicVarWithUnderscore(StringToCharTokens(`X.y...`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}
