package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
	ps "github.com/SSripilaipong/muto/common/parsing"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func TestClass(t *testing.T) {
	t.Run("should parse class", func(t *testing.T) {
		r := Class(StringToCharTokens(`a_bc-'!'.123`))
		expectedResult := st.NewLocalClass(`a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should not parse space", func(t *testing.T) {
		r := Class(StringToCharTokens(`a b`))
		expectedResult := st.NewLocalClass(`a`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` b`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should not parse variable", func(t *testing.T) {
		r := Class(StringToCharTokens(`X`))
		assert.True(t, ps.IsResultErr(r))
	})

	t.Run("should parse symbol", func(t *testing.T) {
		r := Class(StringToCharTokens(`-->.! @`))
		expectedResult := st.NewLocalClass(`-->.!`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` @`))
		assert.Equal(t, tuple.New2(rslt.Value(expectedResult), expectedRemainder), IgnoreLineAndColumnInNewResult(r))
	})

	t.Run("should not match name starting with underscore", func(t *testing.T) {
		r := Class(StringToCharTokens(`_`))
		assert.True(t, ps.IsResultErr(r))
	})

	t.Run("should not match character", func(t *testing.T) {
		r := Class(StringToCharTokens(`'x'`))
		assert.True(t, ps.IsResultErr(r))
	})
}

func TestImportedClass(t *testing.T) {
	r := ImportedClass(StringToCharTokens(`abc.def+xxx`))

	assert.Equal(t, st.NewImportedClass("abc", "def"), ps.ResultValue(r))
	assert.Equal(t, "+xxx", CharactersToString(r.X2()))
}
