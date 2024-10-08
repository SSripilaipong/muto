package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/parser/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestTag(t *testing.T) {
	namedRule := NamedRule()
	t.Run("should parse tag as a child", func(t *testing.T) {
		r := namedRule(base.StringToCharTokens(`f 1 .abc 2`))
		expectedResult := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			st.NewNumber("1"), st.NewTag(`.abc`), st.NewNumber("2"),
		}))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})

	t.Run("should parse tag as a child of a nested named rule", func(t *testing.T) {
		r := namedRule(base.StringToCharTokens(`f (1 .abc 2)`))
		expectedResult := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(st.NewNumber("1"), stPattern.ParamsToFixedParamPart([]stPattern.Param{
				st.NewTag(`.abc`), st.NewNumber("2"),
			})),
		}))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})
}
