package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

func TestString(t *testing.T) {
	r := String(stringToCharTokens(`"abc\n123\""abc`))
	expectedResult := tk.NewString(`"abc\n123\""`)
	expectedRemainder := stringToCharTokens(`abc`)
	assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
}

func TestNumber(t *testing.T) {
	t.Run("should parse positive number", func(t *testing.T) {
		r := Number(stringToCharTokens(`123abc`))
		expectedResult := tk.NewNumber(`123`)
		expectedRemainder := stringToCharTokens(`abc`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should parse positive float number", func(t *testing.T) {
		r := Number(stringToCharTokens(`123.456abc`))
		expectedResult := tk.NewNumber(`123.456`)
		expectedRemainder := stringToCharTokens(`abc`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should parse negative number", func(t *testing.T) {
		r := Number(stringToCharTokens(`-123abc`))
		expectedResult := tk.NewNumber(`-123`)
		expectedRemainder := stringToCharTokens(`abc`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse true", func(t *testing.T) {
		r := Boolean(stringToCharTokens(`trueabc`))
		expectedResult := tk.NewIdentifier(`true`)
		expectedRemainder := stringToCharTokens(`abc`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should parse false", func(t *testing.T) {
		r := Boolean(stringToCharTokens(`falseabc`))
		expectedResult := tk.NewIdentifier(`false`)
		expectedRemainder := stringToCharTokens(`abc`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})
}

func TestFixedVar(t *testing.T) {
	t.Run("should parse single letter variable", func(t *testing.T) {
		r := FixedVar(stringToCharTokens(`X.123`))
		expectedResult := tk.NewIdentifier(`X`)
		expectedRemainder := stringToCharTokens(`.123`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should parse multiple letters variable", func(t *testing.T) {
		r := FixedVar(stringToCharTokens(`XYabc-123!'?-1s.123`))
		expectedResult := tk.NewIdentifier(`XYabc-123!'?-1s`)
		expectedRemainder := stringToCharTokens(`.123`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := FixedVar(stringToCharTokens(`xy`))
		assert.Equal(t, []tuple.Of2[tk.Token, []tk.Token](nil), r)
	})
}

func TestVariadicVar(t *testing.T) {
	t.Run("should parse single letter variadic variable", func(t *testing.T) {
		r := VariadicVar(stringToCharTokens(`X...123`))
		expectedResult := newVariadicVar("X")
		expectedRemainder := stringToCharTokens(`123`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should parse multiple letters variadic variable", func(t *testing.T) {
		r := VariadicVar(stringToCharTokens(`XYabc-123!'?-1s...123`))
		expectedResult := newVariadicVar(`XYabc-123!'?-1s`)
		expectedRemainder := stringToCharTokens(`123`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := VariadicVar(stringToCharTokens(`xy...`))
		assert.Equal(t, []tuple.Of2[VariadicVarNode, []tk.Token](nil), r)
	})

	t.Run("should not parse fixed var", func(t *testing.T) {
		r := VariadicVar(stringToCharTokens(`X.123`))
		assert.Equal(t, []tuple.Of2[VariadicVarNode, []tk.Token](nil), r)
	})

	t.Run("should not parse variadic var with . in the name", func(t *testing.T) {
		r := VariadicVar(stringToCharTokens(`X.y...`))
		assert.Equal(t, []tuple.Of2[VariadicVarNode, []tk.Token](nil), r)
	})
}

func TestClass(t *testing.T) {
	t.Run("should parse class", func(t *testing.T) {
		r := Class(stringToCharTokens(`a_bc-'!'.123`))
		expectedResult := tk.NewIdentifier(`a_bc-'!'`)
		expectedRemainder := stringToCharTokens(`.123`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should not parse space", func(t *testing.T) {
		r := Class(stringToCharTokens(`a b`))
		expectedResult := tk.NewIdentifier(`a`)
		expectedRemainder := stringToCharTokens(` b`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})

	t.Run("should not parse variable", func(t *testing.T) {
		r := Class(stringToCharTokens(`X`))
		assert.Equal(t, []tuple.Of2[tk.Token, []tk.Token](nil), r)
	})

	t.Run("should parse symbol", func(t *testing.T) {
		r := Class(stringToCharTokens(`-->.! @`))
		expectedResult := tk.NewSymbol(`-->.!`)
		expectedRemainder := stringToCharTokens(` @`)
		assert.Equal(t, slc.Pure(tuple.New2(expectedResult, expectedRemainder)), r)
	})
}

func stringToCharTokens(s string) []tk.Token {
	r := make([]tk.Token, len(s))
	for i, x := range []rune(s) {
		r[i] = tk.NewCharacter(x)
	}
	return r
}
