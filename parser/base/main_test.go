package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	st "github.com/SSripilaipong/muto/syntaxtree"
)

func TestString(t *testing.T) {
	r := String(StringToCharTokens(`"abc\n123\""abc`))
	expectedResult := st.NewString(`"abc\n123\""`)
	expectedRemainder := StringToCharTokens(`abc`)
	assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
}

func TestNumber(t *testing.T) {
	t.Run("should parse positive number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123abc`))
		expectedResult := st.NewNumber(`123`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse positive float number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123.456abc`))
		expectedResult := st.NewNumber(`123.456`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse negative number", func(t *testing.T) {
		r := Number(StringToCharTokens(`-123abc`))
		expectedResult := st.NewNumber(`-123`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse true", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`trueabc`))
		expectedResult := st.NewBoolean(`true`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse false", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`falseabc`))
		expectedResult := st.NewBoolean(`false`)
		expectedRemainder := StringToCharTokens(`abc`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})
}

func TestFixedVar(t *testing.T) {
	t.Run("should parse single letter variable", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`X.123`))
		expectedResult := st.NewVariable(`X`)
		expectedRemainder := StringToCharTokens(`.123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should parse multiple letters variable", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`XYabc-123!'?-1s.123`))
		expectedResult := st.NewVariable(`XYabc-123!'?-1s`)
		expectedRemainder := StringToCharTokens(`.123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`xy`))
		assert.Equal(t, EmptyResult[st.Variable](), AsParserResult(r))
	})

	t.Run("should not parse variadic var", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`X...`))
		assert.Equal(t, EmptyResult[st.Variable](), AsParserResult(r))
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
		expectedResult := st.NewClass(`a_bc-'!'`)
		expectedRemainder := StringToCharTokens(`.123`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should not parse space", func(t *testing.T) {
		r := Class(StringToCharTokens(`a b`))
		expectedResult := st.NewClass(`a`)
		expectedRemainder := StringToCharTokens(` b`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})

	t.Run("should not parse variable", func(t *testing.T) {
		r := Class(StringToCharTokens(`X`))
		assert.Equal(t, EmptyResult[st.Class](), AsParserResult(r))
	})

	t.Run("should parse symbol", func(t *testing.T) {
		r := Class(StringToCharTokens(`-->.! @`))
		expectedResult := st.NewClass(`-->.!`)
		expectedRemainder := StringToCharTokens(` @`)
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(r))
	})
}
