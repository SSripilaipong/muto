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
	class := NewBaseModule().GetClass("do")

	t.Run("should become last node", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewOneLayerObject(class,
			base.NewNumberFromString("123"), base.NewNumberFromString("456"),
		))
		assert.Equal(t, base.NewNumberFromString("456"), result)
	})

	t.Run("should mutate with one children", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewOneLayerObject(class,
			base.NewNumberFromString("123"),
		))
		assert.Equal(t, base.NewNumberFromString("123"), result)
	})

	t.Run("should mutate with more than 2 children", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewOneLayerObject(class,
			base.NewNumberFromString("123"), base.NewNumberFromString("456"), base.NewNumberFromString("789"),
		))
		assert.Equal(t, base.NewNumberFromString("789"), result)
	})

	t.Run("should not mutate when no children", func(t *testing.T) {
		result := base.NewOneLayerObject(class).Mutate()
		assert.True(t, result.IsEmpty())
	})
}

func TestControl_match(t *testing.T) {
	module := NewBaseModule()
	class := module.GetClass("match")
	literal := builder.NewLiteralBuilderFactoryWithClassCollection()

	t.Run("should apply first case", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{builder.NewReconstructor( // \A [.ok]
				extractor.NewExactNodeList([]extractor.NodeExtractor{extractor.NewVariable("A")}),
				literal.NewBuilder(stResult.WrapNodeWithObject(st.NewTag(".ok"))),
			)},
			{base.NewNumberFromString("123")}, // 123
		})))
		assert.Equal(t, base.NewOneLayerObject(base.NewTag("ok")), result)
	})

	t.Run("should apply second case", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
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
		assert.Equal(t, base.NewOneLayerObject(base.NewTag("second")), result)
	})
}

func TestControl_ret(t *testing.T) {
	class := NewBaseModule().GetClass("ret")

	t.Run("should apply first case", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewOneLayerObject(class,
			base.NewNumberFromString("123"),
		))
		assert.Equal(t, base.NewNumberFromString("123"), result)
	})
}

func TestControl_compose(t *testing.T) {
	module := NewBaseModule()
	class := module.GetClass("compose")
	stringClass := module.GetClass("string")
	isStringClass := module.GetClass("string?")

	t.Run("should compose 2 functions", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{isStringClass, stringClass},
			{base.NewNumberFromString("123")},
		}))) // (compose string? string) 123
		assert.Equal(t, base.NewBoolean(true), result)
	})
}

func TestControl_curry(t *testing.T) {
	module := NewBaseModule()
	class := module.GetClass("curry")
	addClass := module.GetClass("+")

	t.Run("should curry add function", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{addClass, base.NewNumberFromString("5")},
			{base.NewNumberFromString("20")},
		}))) // (curry + 5) 20
		assert.Equal(t, base.NewNumberFromString("25"), result)
	})
}

func TestControl_with(t *testing.T) {
	module := NewBaseModule()
	class := module.GetClass("with")
	literal := builder.NewLiteralBuilderFactoryWithClassCollection()

	t.Run("should match variables", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{base.NewString("abc"), base.NewTag("xxx")},
			{builder.NewReconstructor(
				extractor.NewExactNodeList([]extractor.NodeExtractor{extractor.NewVariable("A"), extractor.NewTag("xxx")}),
				literal.NewBuilder(stResult.NewObject(st.NewTag(".ok"), []stResult.Param{st.NewVariable("A")})),
			)},
		}))) // (with "abc" .xxx) \A .xxx [.ok A]
		assert.Equal(t, base.NewOneLayerObject(base.NewTag("ok"), base.NewString("abc")), result)
	})
}

func TestControl_use(t *testing.T) {
	module := NewBaseModule()
	class := module.GetClass("use")

	t.Run("should match variables", func(t *testing.T) {
		result := base.MutateUntilTerminated(base.NewCompoundObject(class, base.NewParamChain([][]base.Node{
			{base.NewTag("ok")},
			{base.NewConventionalList(base.NewRune('a'), base.NewRune('b'))},
		}))) // (use .ok) ($ 'a' 'b')
		assert.Equal(t, base.NewOneLayerObject(base.NewTag("ok"), base.NewRune('a'), base.NewRune('b')), result)
	})
}
