package result

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestTag(t *testing.T) {
	node := SimplifiedNode()
	t.Run("should parse tag as result", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`.abc`))
		expectedResult := stResult.ToSimplifiedNode(stResult.SingleNodeToNakedObject(base.NewTag(`.abc`)))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(r))
	})

	t.Run("should parse tag as nested head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`1 (.abc 2)`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			base.NewNumber("1"),
			stResult.ParamsToFixedParamPart([]stResult.Param{
				stResult.NewObject(
					base.NewTag(`.abc`),
					stResult.ParamsToFixedParamPart([]stResult.Param{base.NewNumber("2")}),
				),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse tag as nested child", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`1 (2 .abc)`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			base.NewNumber("1"),
			stResult.ParamsToFixedParamPart([]stResult.Param{
				stResult.NewObject(
					base.NewNumber("2"),
					stResult.ParamsToFixedParamPart([]stResult.Param{base.NewTag(`.abc`)}),
				),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_Structure(t *testing.T) {
	node := SimplifiedNode()
	emptyStructure := stResult.NewStructure([]stResult.StructureRecord{})

	t.Run("should parse structure as result", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("{}"))
		expectedResult := stResult.ToSimplifiedNode(stResult.SingleNodeToNakedObject(emptyStructure))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(r))
	})

	t.Run("should parse structure as object param", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f {}"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{emptyStructure})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse structure as object head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("{} f"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(emptyStructure, stResult.ParamsToFixedParamPart([]stResult.Param{base.NewClass("f")})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse structure as nested object param", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f (g {})"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
			stResult.NewObject(base.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{emptyStructure})),
		})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_nestedHead(t *testing.T) {
	node := SimplifiedNode()

	t.Run("should parse object as a head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`(p) "a"`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			stResult.NewObject(base.NewClass("p"), stResult.ParamsToFixedParamPart([]stResult.Param{})),
			stResult.ParamsToFixedParamPart([]stResult.Param{base.NewString(`"a"`)}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse double-nested head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`((p)) "a"`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			stResult.NewObject(stResult.NewObject(base.NewClass("p"), stResult.ParamsToFixedParamPart([]stResult.Param{})), stResult.ParamsToFixedParamPart([]stResult.Param{})),
			stResult.ParamsToFixedParamPart([]stResult.Param{base.NewString(`"a"`)}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_nestedObject(t *testing.T) {
	node := SimplifiedNode()

	t.Run("should parse nested object with no params", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`g (f)`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(base.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{
			stResult.NewObject(base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{})),
		})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_objectMultilines(t *testing.T) {
	node := SimplifiedNode()

	t.Run("should allow naked multiline object with uniform indent block", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f\n  g\n  1\n  2"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				base.NewClass("g"), base.NewNumber("1"), base.NewNumber(`2`),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow naked multiline object with uniform indent block starting with single-line param", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f  g\n  1\n  2"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				base.NewClass("g"), base.NewNumber("1"), base.NewNumber(`2`),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object inside parentheses", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("(\nf \n  g\n\n 1 2 )"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewObject(
			base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				base.NewClass("g"), base.NewNumber("1"), base.NewNumber(`2`),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object in param part", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("a (\nf g\n\n 1 2  )"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(base.NewClass("a"), stResult.ParamsToFixedParamPart([]stResult.Param{
			stResult.NewObject(base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				base.NewClass("g"), base.NewNumber("1"), base.NewNumber(`2`),
			})),
		})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object as an object head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("(\n\nf\ng\n1\n2  \n) x"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			stResult.NewObject(base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				base.NewClass("g"), base.NewNumber("1"), base.NewNumber(`2`),
			})),
			stResult.ParamsToFixedParamPart([]stResult.Param{base.NewClass("x")}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object as structure's value", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("{.a: ( \nf\ng\n1\n2  )}"))
		expectedResult := stResult.ToSimplifiedNode(stResult.SingleNodeToNakedObject(stResult.NewStructure(slc.Pure(stResult.NewStructureRecord(
			base.NewTag(".a"),
			stResult.NewObject(base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				base.NewClass("g"), base.NewNumber("1"), base.NewNumber(`2`),
			})),
		)))))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}
