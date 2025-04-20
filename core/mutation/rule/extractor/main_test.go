package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestNew_ObjectWithValueHead(t *testing.T) {
	t.Run("should match boolean as a nested object head", func(t *testing.T) {
		// pattern: f (true 456)
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(stBase.NewBoolean("true"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{stBase.NewNumber("456")})),
		}))
		// obj: f (true 456)
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewOneLayerObject(base.NewBoolean(true), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(stBase.NewString(`"a"`), stPattern.PatternsToFixedParamPart([]stBase.Pattern{stBase.NewNumber("456")})),
		}))
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewOneLayerObject(base.NewString(`a`), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match number as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(stBase.NewNumber("123"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{stBase.NewNumber("456")})),
		}))
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewOneLayerObject(base.NewNumberFromString("123"), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})
}

func TestTag(t *testing.T) {
	t.Run("should match a tag child", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{stBase.NewTag("abc")}))
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewTag("abc")})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should not match a non-tag child", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{stBase.NewTag("abc")}))
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewClass("abc")})

		assert.False(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested head", func(t *testing.T) {
		// pattern: f (.abc 1 2)
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(stBase.NewTag("abc"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
				stBase.NewNumber("1"), stBase.NewNumber("2"),
			})),
		}))
		// obj: f (.abc 1 2)
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewOneLayerObject(base.NewTag("abc"), []base.Node{base.NewNumberFromString("1"), base.NewNumberFromString("2")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested param", func(t *testing.T) {
		// pattern: f (g 1 .abc 2)
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(stBase.NewClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
				stBase.NewNumber("1"), stBase.NewTag("abc"), stBase.NewNumber("2"),
			})),
		}))
		// obj: f (g 1 .abc 2)
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewOneLayerObject(base.NewClass("g"), []base.Node{base.NewNumberFromString("1"), base.NewTag("abc"), base.NewNumberFromString("2")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})
}

func TestNestedObject(t *testing.T) {
	t.Run("should not match leaf object with simple node", func(t *testing.T) {
		// pattern: f (g)
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(stBase.NewClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{})),
		}))
		// obj: f g
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewClass("g")})

		assert.True(t, New(pattern).Extract(obj).IsEmpty())
	})
}

func TestVariadicParam(t *testing.T) {
	t.Run("should match nested variadic param with size 0", func(t *testing.T) {
		// pattern: g (f Xs...)
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(stBase.NewClass("f"),
				stPattern.NewLeftVariadicParamPart("Xs", stPattern.FixedParamPart{})),
		}))
		// obj: g (f)
		obj := base.NewNamedOneLayerObject("g", []base.Node{base.NewNamedOneLayerObject("f", nil)})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should ignore underscore variadic variable", func(t *testing.T) {
		// pattern: g (f _Bc...)
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(stBase.NewClass("f"),
				stPattern.NewLeftVariadicParamPart("_Bc", stPattern.FixedParamPart{})),
		}))
		// obj: g (f x y)
		obj := base.NewNamedOneLayerObject("g", []base.Node{
			base.NewNamedOneLayerObject("f", []base.Node{base.NewClass("x"), base.NewClass("y")}),
		})

		p := New(pattern).Extract(obj)
		assert.True(t, p.IsNotEmpty() && p.Value().VariadicVarValue("_Bc").IsEmpty())
	})
}

func TestVariableParam(t *testing.T) {
	t.Run("should ignore underscore variable", func(t *testing.T) {
		// pattern: g (f _Bc)
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(stBase.NewClass("f"),
				stPattern.NewLeftVariadicParamPart("_Bc", stPattern.FixedParamPart{})),
		}))
		// obj: g (f x)
		obj := base.NewNamedOneLayerObject("g", []base.Node{
			base.NewNamedOneLayerObject("f", []base.Node{base.NewClass("x"), base.NewClass("y")}),
		})

		p := New(pattern).Extract(obj)
		assert.True(t, p.IsNotEmpty() && p.Value().VariableValue("_Bc").IsEmpty())
	})
}

func TestRemainingChildren(t *testing.T) {
	t.Run("should extract remaining children", func(t *testing.T) {
		// pattern: f 1
		pattern := stPattern.NewDeterminantObject(stBase.NewClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stBase.NewNumber("1"),
		}))
		// obj: f 1 2 3
		obj := base.NewNamedOneLayerObject("f", []base.Node{base.NewNumberFromString("1"), base.NewNumberFromString("2"), base.NewNumberFromString("3")})

		assert.Equal(t, []base.Node{base.NewNumberFromString("2"), base.NewNumberFromString("3")}, New(pattern).Extract(obj).Value().RemainingParamChain().All()[0])
	})
}
