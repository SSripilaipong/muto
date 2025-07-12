package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestFixedVar(t *testing.T) {
	t.Run("should parse single letter variable", func(t *testing.T) {
		r := FixedVarWithUnderscore(StringToCharTokens(`X.123`))
		expectedResult := syntaxtree.NewVariable(`X`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse multiple letters variable", func(t *testing.T) {
		r := FixedVarWithUnderscore(StringToCharTokens(`XYabc-123!'?-1s.123`))
		expectedResult := syntaxtree.NewVariable(`XYabc-123!'?-1s`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := FixedVarWithUnderscore(StringToCharTokens(`xy`))
		assert.Equal(t, EmptyResult[syntaxtree.Variable](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse variadic var", func(t *testing.T) {
		r := FixedVarWithUnderscore(IgnoreLineAndColumn(StringToCharTokens(`X...`)))
		assert.Equal(t, EmptyResult[syntaxtree.Variable](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}
