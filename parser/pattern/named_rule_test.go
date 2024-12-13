package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/parser/base"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestTag(t *testing.T) {
	namedRule := NamedRule()
	t.Run("should parse tag as a child", func(t *testing.T) {
		r := namedRule(base.StringToCharTokens(`f 1 .abc 2`))
		expectedResult := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stBase.NewNumber("1"), stBase.NewTag(`.abc`), stBase.NewNumber("2"),
		}))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})

	t.Run("should parse tag as a child of a nested named rule", func(t *testing.T) {
		r := namedRule(base.StringToCharTokens(`f (1 .abc 2)`))
		expectedResult := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(stBase.NewNumber("1"), stPattern.ParamsToFixedParamPart([]stPattern.Param{
				stBase.NewTag(`.abc`), stBase.NewNumber("2"),
			})),
		}))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})
}
