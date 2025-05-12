package result

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/slc"
	psBase "github.com/SSripilaipong/muto/parser/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestTag(t *testing.T) {
	node := SimplifiedNode()
	t.Run("should parse tag as result", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`.abc`))
		expectedResult := stResult.ToSimplifiedNode(stResult.SingleNodeToNakedObject(st.NewTag(`.abc`)))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(r))
	})

	t.Run("should parse tag as nested head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`1 (.abc 2)`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			st.NewNumber("1"),
			stResult.ParamsToFixedParamPart([]stResult.Param{
				stResult.NewObject(
					st.NewTag(`.abc`),
					stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("2")}),
				),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse tag as nested child", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`1 (2 .abc)`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			st.NewNumber("1"),
			stResult.ParamsToFixedParamPart([]stResult.Param{
				stResult.NewObject(
					st.NewNumber("2"),
					stResult.ParamsToFixedParamPart([]stResult.Param{st.NewTag(`.abc`)}),
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
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{emptyStructure})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse structure as object head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("{} f"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(emptyStructure, stResult.ParamsToFixedParamPart([]stResult.Param{st.NewClass("f")})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse structure as nested object param", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f (g {})"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
			stResult.NewObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{emptyStructure})),
		})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_nestedHead(t *testing.T) {
	node := SimplifiedNode()

	t.Run("should parse object as a head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`(p) "a"`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			stResult.NewObject(st.NewClass("p"), stResult.ParamsToFixedParamPart([]stResult.Param{})),
			stResult.ParamsToFixedParamPart([]stResult.Param{st.NewString(`"a"`)}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should parse double-nested head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`((p)) "a"`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			stResult.NewObject(stResult.NewObject(st.NewClass("p"), stResult.ParamsToFixedParamPart([]stResult.Param{})), stResult.ParamsToFixedParamPart([]stResult.Param{})),
			stResult.ParamsToFixedParamPart([]stResult.Param{st.NewString(`"a"`)}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_nestedObject(t *testing.T) {
	node := SimplifiedNode()

	t.Run("should parse nested object with no params", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`g (f)`))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{
			stResult.NewObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{})),
		})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_objectMultilines(t *testing.T) {
	node := SimplifiedNode()

	t.Run("should allow naked multiline object with uniform indent block", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f\n  g\n  1\n  2"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				st.NewClass("g"), st.NewNumber("1"), st.NewNumber(`2`),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow naked multiline object with uniform indent block starting with single-line param", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("f  g\n  1\n  2"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				st.NewClass("g"), st.NewNumber("1"), st.NewNumber(`2`),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object inside parentheses", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("(\nf \n  g\n\n 1 2 )"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewObject(
			st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				st.NewClass("g"), st.NewNumber("1"), st.NewNumber(`2`),
			}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object in param part", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("a (\nf g\n\n 1 2  )"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(st.NewClass("a"), stResult.ParamsToFixedParamPart([]stResult.Param{
			stResult.NewObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				st.NewClass("g"), st.NewNumber("1"), st.NewNumber(`2`),
			})),
		})))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object as an object head", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("(\n\nf\ng\n1\n2  \n) x"))
		expectedResult := stResult.ToSimplifiedNode(stResult.NewNakedObject(
			stResult.NewObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				st.NewClass("g"), st.NewNumber("1"), st.NewNumber(`2`),
			})),
			stResult.ParamsToFixedParamPart([]stResult.Param{st.NewClass("x")}),
		))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})

	t.Run("should allow multiline object as structure's value", func(t *testing.T) {
		r := node(psBase.StringToCharTokens("{.a: ( \nf\ng\n1\n2  )}"))
		expectedResult := stResult.ToSimplifiedNode(stResult.SingleNodeToNakedObject(stResult.NewStructure(slc.Pure(stResult.NewStructureRecord(
			st.NewTag(".a"),
			stResult.NewObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{
				st.NewClass("g"), st.NewNumber("1"), st.NewNumber(`2`),
			})),
		)))))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}

func TestNode_Reconstructor(t *testing.T) {
	node := SimplifiedNode()

	t.Run("should parse identity reconstructor", func(t *testing.T) {
		r := node(psBase.StringToCharTokens(`\X[ret X]`))
		expectedResult := stResult.ToSimplifiedNode(stResult.SingleNodeToNakedObject(stResult.NewReconstructor(
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{st.NewVariable("X")}),
			stResult.NewObject(st.NewClass("ret"), stResult.ParamsToFixedParamPart([]stResult.Param{st.NewVariable("X")})),
		)))
		assert.Equal(t, psBase.SingleResult(expectedResult, []psBase.Character{}), psBase.AsParserResult(parsing.FilterSuccess(r)))
	})
}
