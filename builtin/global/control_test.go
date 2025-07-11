package global

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestControl_do(t *testing.T) {
	class := NewModuleForStdio().GetOrCreateClass("do")
	t.Run("should become last node", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			base.NewNumberFromString("123"), base.NewNumberFromString("456"),
		}))
		assert.Equal(t, base.NewNumberFromString("456"), result)
	})

	t.Run("should mutate with one children", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			base.NewNumberFromString("123"),
		}))
		assert.Equal(t, base.NewNumberFromString("123"), result)
	})

	t.Run("should mutate with more than 2 children", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			base.NewNumberFromString("123"), base.NewNumberFromString("456"), base.NewNumberFromString("789"),
		}))
		assert.Equal(t, base.NewNumberFromString("789"), result)
	})

	t.Run("should not mutate when no children", func(t *testing.T) {
		result := base.NewOneLayerObject(class, []base.Node{}).Mutate()
		assert.True(t, result.IsEmpty())
	})
}

func TestControl_match(t *testing.T) {
	class := NewModuleForStdio().GetOrCreateClass("match")
	literal := builder.NewLiteralBuilderFactoryWithClassCollection(builder.NewStaticClassCollection())

	t.Run("should apply first case", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{builder.NewReconstructor( // \A [.ok]
				extractor.NewExactNodeList([]extractor.NodeExtractor{extractor.NewVariable("A")}),
				literal.NewBuilder(stResult.WrapNodeWithObject(st.NewTag(".ok"))),
			)},
			{base.NewNumberFromString("123")}, // 123
		})))
		assert.Equal(t, base.WrapWithObject(base.NewTag("ok")), result)
	})

	t.Run("should apply second case", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{
				builder.NewReconstructor( // \"not this one" [.first]
					extractor.NewExactNodeList([]extractor.NodeExtractor{extractor.NewString("not this one")}),
					literal.NewBuilder(stResult.WrapNodeWithObject(st.NewTag(".first"))),
				),
				builder.NewReconstructor( // \"yeah" [.second]
					extractor.NewExactNodeList([]extractor.NodeExtractor{extractor.NewString("yeah")}),
					literal.NewBuilder(stResult.WrapNodeWithObject(st.NewTag(".second"))),
				),
			},
			{base.NewString("yeah")}, // "yeah"
		})))
		assert.Equal(t, base.WrapWithObject(base.NewTag("second")), result)
	})
}

func TestControl_ret(t *testing.T) {
	class := NewModuleForStdio().GetOrCreateClass("ret")

	t.Run("should apply first case", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			base.NewNumberFromString("123"),
		}))
		assert.Equal(t, base.NewNumberFromString("123"), result)
	})
}

func TestControl_compose(t *testing.T) {
	module := NewModuleForStdio()
	class := module.GetOrCreateClass("compose")
	stringClass := module.GetOrCreateClass("string")
	isStringClass := module.GetOrCreateClass("string?")

	t.Run("should compose 2 functions", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{isStringClass, stringClass},
			{base.NewNumberFromString("123")},
		}))) // (compose string? string) 123
		assert.Equal(t, base.NewBoolean(true), result)
	})
}

func TestControl_curry(t *testing.T) {
	module := NewModuleForStdio()
	class := module.GetOrCreateClass("curry")
	addClass := module.GetOrCreateClass("+")

	t.Run("should curry add function", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{addClass, base.NewNumberFromString("5")},
			{base.NewNumberFromString("20")},
		}))) // (curry + 5) 20
		assert.Equal(t, base.NewNumberFromString("25"), result)
	})
}
