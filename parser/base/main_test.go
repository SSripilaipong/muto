package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

func TestString(t *testing.T) {
	r := String(StringToCharTokens(`"abc\n123\""abc`))
	expectedResult := stBase.NewString(`"abc\n123\""`)
	expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
	assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
}

func TestNumber(t *testing.T) {
	t.Run("should parse positive number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123abc`))
		expectedResult := stBase.NewNumber(`123`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse positive float number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123.456abc`))
		expectedResult := stBase.NewNumber(`123.456`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse negative number", func(t *testing.T) {
		r := Number(StringToCharTokens(`-123abc`))
		expectedResult := stBase.NewNumber(`-123`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse true", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`trueabc`))
		expectedResult := stBase.NewBoolean(`true`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse false", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`falseabc`))
		expectedResult := stBase.NewBoolean(`false`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestFixedVar(t *testing.T) {
	t.Run("should parse single letter variable", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`X.123`))
		expectedResult := stBase.NewVariable(`X`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse multiple letters variable", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`XYabc-123!'?-1s.123`))
		expectedResult := stBase.NewVariable(`XYabc-123!'?-1s`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := FixedVar(StringToCharTokens(`xy`))
		assert.Equal(t, EmptyResult[stBase.Variable](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse variadic var", func(t *testing.T) {
		r := FixedVar(IgnoreLineAndColumn(StringToCharTokens(`X...`)))
		assert.Equal(t, EmptyResult[stBase.Variable](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestVariadicVar(t *testing.T) {
	t.Run("should parse single letter variadic variable", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`X...123`))
		expectedResult := newVariadicVar("X")
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse multiple letters variadic variable", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`XYabc-123!'?-1s...123`))
		expectedResult := newVariadicVar(`XYabc-123!'?-1s`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse identifier starting with small case", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`xy...`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse fixed var", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`X.123`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse variadic var with . in the name", func(t *testing.T) {
		r := VariadicVar(StringToCharTokens(`X.y...`))
		assert.Equal(t, EmptyResult[VariadicVarNode](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestClass(t *testing.T) {
	t.Run("should parse class", func(t *testing.T) {
		r := Class(StringToCharTokens(`a_bc-'!'.123`))
		expectedResult := stBase.NewClass(`a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse space", func(t *testing.T) {
		r := Class(StringToCharTokens(`a b`))
		expectedResult := stBase.NewClass(`a`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` b`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse variable", func(t *testing.T) {
		r := Class(StringToCharTokens(`X`))
		assert.Equal(t, EmptyResult[stBase.Class](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse symbol", func(t *testing.T) {
		r := Class(StringToCharTokens(`-->.! @`))
		expectedResult := stBase.NewClass(`-->.!`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` @`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestTag(t *testing.T) {
	t.Run("should parse tag like class with dot prefix", func(t *testing.T) {
		r := Tag(StringToCharTokens(`.a_bc-'!'.123`))
		expectedResult := stBase.NewTag(`.a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse class", func(t *testing.T) {
		r := Tag(StringToCharTokens(`abc`))
		assert.Equal(t, EmptyResult[stBase.Tag](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}
