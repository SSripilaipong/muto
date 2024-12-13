package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestNew_ObjectWithValueHead(t *testing.T) {
	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(stBase.NewBoolean("true"), stPattern.ParamsToFixedParamPart([]stPattern.Param{stBase.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewBoolean(true), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(stBase.NewString(`"a"`), stPattern.ParamsToFixedParamPart([]stPattern.Param{stBase.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewString(`a`), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match number as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(stBase.NewNumber("123"), stPattern.ParamsToFixedParamPart([]stPattern.Param{stBase.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewNumberFromString("123"), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})
}

func TestTag(t *testing.T) {
	t.Run("should match a tag child", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{stBase.NewTag("abc")}))
		obj := base.NewNamedObject("f", []base.Node{base.NewTag("abc")})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should not match a non-tag child", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{stBase.NewTag("abc")}))
		obj := base.NewNamedObject("f", []base.Node{base.NewClass("abc")})

		assert.False(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewAnonymousRule(stBase.NewTag("abc"), stPattern.ParamsToFixedParamPart([]stPattern.Param{
				stBase.NewNumber("1"), stBase.NewNumber("2"),
			})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewTag("abc"), []base.Node{base.NewNumberFromString("1"), base.NewNumberFromString("2")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested head", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewNamedRule("g", stPattern.ParamsToFixedParamPart([]stPattern.Param{
				stBase.NewNumber("1"), stBase.NewTag("abc"), stBase.NewNumber("2"),
			})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewClass("g"), []base.Node{base.NewNumberFromString("1"), base.NewTag("abc"), base.NewNumberFromString("2")})})

		assert.True(t, New(pattern).Extract(obj).IsNotEmpty())
	})
}

func TestNestedObject(t *testing.T) {
	t.Run("should not match leaf object with simple node", func(t *testing.T) {
		pattern := stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
			stPattern.NewNamedRule("g", stPattern.ParamsToFixedParamPart([]stPattern.Param{})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewClass("g")})

		assert.True(t, New(pattern).Extract(obj).IsEmpty())
	})
}
