package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

func TestString(t *testing.T) {
	r := String(StringToCharTokens(`"abc\n123\""abc`))
	expectedResult := tk.NewString(`"abc\n123\""`)
	expectedRemainder := StringToCharTokens(`abc`)
	assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
}

func TestNumber(t *testing.T) {
	t.Run("should parse positive number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123abc`))
		expectedResult := tk.NewNumber(`123`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse positive float number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123.456abc`))
		expectedResult := tk.NewNumber(`123.456`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse negative number", func(t *testing.T) {
		r := Number(StringToCharTokens(`-123abc`))
		expectedResult := tk.NewNumber(`-123`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse true", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`trueabc`))
		expectedResult := tk.NewIdentifier(`true`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse false", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`falseabc`))
		expectedResult := tk.NewIdentifier(`false`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})
}

func TestFixedVar(t *testing.T) {
	t.Run("should parse single letter variable", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`X.123`))
		expectedResult := tk.NewIdentifier(`X`)
		expectedRemainder := StringToCharTokens(`.123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse multiple letters variable", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`XYabc-123!'?-1s.123`))
		expectedResult := tk.NewIdentifier(`XYabc-123!'?-1s`)
		expectedRemainder := StringToCharTokens(`.123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`xy`))
		assert.Equal(t, EmptyResult[tk.Token](), AsParserResult(r))
	})
}

func TestVariadicVar(t *testing.T) {
	t.Run("should parse single letter variadic variable", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`X...123`))
		expectedResult := newVariadicVar("X")
		expectedRemainder := StringToCharTokens(`123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse multiple letters variadic variable", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`XYabc-123!'?-1s...123`))
		expectedResult := newVariadicVar(`XYabc-123!'?-1s`)
		expectedRemainder := StringToCharTokens(`123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`xy...`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(r))
	})

	t.Run("should not parse fixed var", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`X.123`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(r))
	})

	t.Run("should not parse variadic var with . in the name", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`X.y...`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(r))
	})
}

func TestClass(t *testing.T) {
	t.Run("should parse class", func(t *testing.T) {
		r := Class(StringToCharTokens(`a_bc-'!'.123`))
		expectedResult := tk.NewIdentifier(`a_bc-'!'`)
		expectedRemainder := StringToCharTokens(`.123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should not parse space", func(t *testing.T) {
		r := Class(StringToCharTokens(`a b`))
		expectedResult := tk.NewIdentifier(`a`)
		expectedRemainder := StringToCharTokens(` b`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should not parse variable", func(t *testing.T) {
		r := Class(StringToCharTokens(`X`))
		assert.Equal(t, EmptyResult[tk.Token](), AsParserResult(r))
	})

	t.Run("should parse symbol", func(t *testing.T) {
		r := Class(StringToCharTokens(`-->.! @`))
		expectedResult := tk.NewSymbol(`-->.!`)
		expectedRemainder := StringToCharTokens(` @`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})
}
