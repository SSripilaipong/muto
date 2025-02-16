package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/parser/base"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestTag(t *testing.T) {
	namedRule := Pattern()
	t.Run("should parse tag as a child", func(t *testing.T) {
		r := namedRule(base.StringToCharTokens(`f 1 .abc 2`))
		expectedResult := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stBase.PatternParam{
			stBase.NewNumber("1"), stBase.NewTag(`.abc`), stBase.NewNumber("2"),
		}))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})

	t.Run("should parse tag as a child of a nested named rule", func(t *testing.T) {
		r := namedRule(base.StringToCharTokens(`f (1 .abc 2)`))
		expectedResult := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stBase.PatternParam{
			stPattern.NewAnonymousRule(stBase.NewNumber("1"), stPattern.ParamsToFixedParamPart([]stBase.PatternParam{
				stBase.NewTag(`.abc`), stBase.NewNumber("2"),
			})),
		}))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})
}

func TestNamedRule_variadic(t *testing.T) {
	namedRule := Pattern()

	t.Run("should parse nested variadic param", func(t *testing.T) {
		r := namedRule(base.StringToCharTokens(`g (f Xs...)`))
		expectedResult := stPattern.NewNamedRule("g", stPattern.ParamsToFixedParamPart([]stBase.PatternParam{
			stPattern.NewNamedRule("f",
				stPattern.NewLeftVariadicParamPart("Xs", stPattern.FixedParamPart{})),
		}))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})
}
