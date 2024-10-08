package result

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/parser/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestTag(t *testing.T) {
	t.Run("should parse tag as result", func(t *testing.T) {
		r := Node(base.StringToCharTokens(`.abc`))
		expectedResult := stResult.ToNode(st.NewTag(`.abc`))
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(r))
	})

	t.Run("should parse tag as nested head", func(t *testing.T) {
		r := Node(base.StringToCharTokens(`1 (.abc 2)`))
		expectedResult := stResult.ToNode(
			stResult.NewObject(
				st.NewNumber("1"),
				stResult.ParamsToParamPart([]stResult.Param{
					stResult.NewObject(
						st.NewTag(`.abc`),
						stResult.ParamsToParamPart([]stResult.Param{st.NewNumber("2")}),
					),
				}),
			),
		)
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse tag as nested child", func(t *testing.T) {
		r := Node(base.StringToCharTokens(`1 (2 .abc)`))
		expectedResult := stResult.ToNode(
			stResult.NewObject(
				st.NewNumber("1"),
				stResult.ParamsToParamPart([]stResult.Param{
					stResult.NewObject(
						st.NewNumber("2"),
						stResult.ParamsToParamPart([]stResult.Param{st.NewTag(`.abc`)}),
					),
				}),
			),
		)
		assert.Equal(t, base.SingleResult(expectedResult, []base.Character{}), base.AsParserResult(parsing.FilterSuccess(r)))
	})
}
