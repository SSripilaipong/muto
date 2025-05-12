package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/syntaxtree"
)

func TestString(t *testing.T) {
	r := String(StringToCharTokens(`"abc\n123\""abc`))
	expectedResult := syntaxtree.NewString(`"abc\n123\""`)
	expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
	assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
}

func TestNumber(t *testing.T) {
	t.Run("should parse positive number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123abc`))
		expectedResult := syntaxtree.NewNumber(`123`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse positive float number", func(t *testing.T) {
		r := Number(StringToCharTokens(`123.456abc`))
		expectedResult := syntaxtree.NewNumber(`123.456`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse negative number", func(t *testing.T) {
		r := Number(StringToCharTokens(`-123abc`))
		expectedResult := syntaxtree.NewNumber(`-123`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse true", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`trueabc`))
		expectedResult := syntaxtree.NewBoolean(`true`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse false", func(t *testing.T) {
		r := Boolean(StringToCharTokens(`falseabc`))
		expectedResult := syntaxtree.NewBoolean(`false`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`abc`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

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

func TestClass(t *testing.T) {
	t.Run("should parse class", func(t *testing.T) {
		r := Class(StringToCharTokens(`a_bc-'!'.123`))
		expectedResult := syntaxtree.NewClass(`a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse space", func(t *testing.T) {
		r := Class(StringToCharTokens(`a b`))
		expectedResult := syntaxtree.NewClass(`a`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` b`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse variable", func(t *testing.T) {
		r := Class(StringToCharTokens(`X`))
		assert.Equal(t, EmptyResult[syntaxtree.Class](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should parse symbol", func(t *testing.T) {
		r := Class(StringToCharTokens(`-->.! @`))
		expectedResult := syntaxtree.NewClass(`-->.!`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(` @`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not match name starting with underscore", func(t *testing.T) {
		r := Class(StringToCharTokens(`_`))
		assert.Equal(t, EmptyResult[syntaxtree.Class](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}

func TestTag(t *testing.T) {
	t.Run("should parse tag like class with dot prefix", func(t *testing.T) {
		r := Tag(StringToCharTokens(`.a_bc-'!'.123`))
		expectedResult := syntaxtree.NewTag(`.a_bc-'!'`)
		expectedRemainder := IgnoreLineAndColumn(StringToCharTokens(`.123`))
		assert.Equal(t, SingleResult(expectedResult, expectedRemainder), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})

	t.Run("should not parse class", func(t *testing.T) {
		r := Tag(StringToCharTokens(`abc`))
		assert.Equal(t, EmptyResult[syntaxtree.Tag](), AsParserResult(IgnoreLineAndColumnInResult(r)))
	})
}
