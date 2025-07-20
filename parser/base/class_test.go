package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	st "github.com/SSripilaipong/muto/syntaxtree"
)

func TestClass(t *testing.T) {
	t.Run("should parse class", func(t *testing.T) {
		r := Class(StringToCharTokens(`a_bc-'!'.123`))
		expectedResult := st.NewLocalClass(`a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse space", func(t *testing.T) {
		r := Class(StringToCharTokens(`a b`))
		expectedResult := st.NewLocalClass(`a`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` b`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse variable", func(t *testing.T) {
		r := Class(StringToCharTokens(`X`))
		assert.Equal(t, EmptyResult[st.LocalClass](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse symbol", func(t *testing.T) {
		r := Class(StringToCharTokens(`-->.! @`))
		expectedResult := st.NewLocalClass(`-->.!`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` @`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not match name starting with underscore", func(t *testing.T) {
		r := Class(StringToCharTokens(`_`))
		assert.Equal(t, EmptyResult[st.LocalClass](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not match character", func(t *testing.T) {
		r := Class(StringToCharTokens(`'x'`))
		assert.Equal(t, EmptyResult[st.LocalClass](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestImportedClass(t *testing.T) {
	result := ImportedClass(StringToCharTokens(`abc.def+xxx`))

	require.Len(t, result, 1)
	r := result[0]

	assert.Equal(t, st.NewImportedClass("abc", "def"), r.X1())
	assert.Equal(t, "+xxx", CharactersToString(r.X2()))
}
