package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestChar(t *testing.T) {
	t.Run("should parse non-escaped rune", func(t *testing.T) {
		r := Rune(StringToCharTokens(`'x'abc`))
		expectedResult := syntaxtree.NewRune(`'x'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse escaped rune", func(t *testing.T) {
		r := Rune(StringToCharTokens(`'\''abc`))
		expectedResult := syntaxtree.NewRune(`'\''`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse unicode", func(t *testing.T) {
		r := Rune(StringToCharTokens(`'μ'abc`))
		expectedResult := syntaxtree.NewRune(`'μ'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}
