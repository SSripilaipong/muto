package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestNew_ObjectWithValueHead(t *testing.T) {
	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(st.NewBoolean("true"), stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewBoolean(true), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})

	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(st.NewString(`"a"`), stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewString(`a`), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})

	t.Run("should match number as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(st.NewNumber("123"), stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewNumberFromString("123"), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})
}

func TestTag(t *testing.T) {
	t.Run("should match a tag child", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewTag("abc")}))
		obj := base.NewNamedObject("f", []base.Node{base.NewTag("abc")})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})

	t.Run("should not match a non-tag child", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewTag("abc")}))
		obj := base.NewNamedObject("f", []base.Node{base.NewClass("abc")})

		assert.False(t, New(pattern)(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(st.NewTag("abc"), stPattern.ParamsToFixedParamPart([]stPattern.Param{
				st.NewNumber("1"), st.NewNumber("2"),
			})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewTag("abc"), []base.Node{base.NewNumberFromString("1"), base.NewNumberFromString("2")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewNamedRule("g", stPattern.ParamsToFixedParamPart([]stPattern.Param{
				st.NewNumber("1"), st.NewTag("abc"), st.NewNumber("2"),
			})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewClass("g"), []base.Node{base.NewNumberFromString("1"), base.NewTag("abc"), base.NewNumberFromString("2")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})
}
