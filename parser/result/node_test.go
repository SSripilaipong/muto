package result

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestTag(t *testing.T) {
	node := Node()
	t.Run("should parse tag as result", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`.abc`))
		expectedResult := stResult.ToNode(st.NewTag(`.abc`))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(r))
	})

	t.Run("should parse tag as nested head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`1 (.abc 2)`))
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
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse tag as nested child", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`1 (2 .abc)`))
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
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_Structure(t *testing.T) {
	node := Node()
	emptyStructure := stResult.NewStructure([]stResult.StructureRecord{})

	t.Run("should parse structure as result", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("{}"))
		expectedResult := stResult.ToNode(emptyStructure)
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(r))
	})

	t.Run("should parse structure as object param", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f {}"))
		expectedResult := stResult.ToNode(stResult.NewObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{emptyStructure})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse structure as object head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("{} f"))
		expectedResult := stResult.ToNode(stResult.NewObject(emptyStructure, stResult.ParamsToFixedParamPart([]stResult.Param{st.NewClass("f")})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse structure as nested object param", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f (g {})"))
		expectedResult := stResult.ToNode(stResult.NewObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
			stResult.NewObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{emptyStructure})),
		})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}
