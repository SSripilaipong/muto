package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func TestNew_ObjectWithValueHead(t *testing.T) {
	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := st.NewNamedRulePattern("f", st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{
			st.NewAnonymousRulePattern(st.NewBoolean("true"), st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{st.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewBoolean(true), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})

	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := st.NewNamedRulePattern("f", st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{
			st.NewAnonymousRulePattern(st.NewString(`"a"`), st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{st.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewString(`a`), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})

	t.Run("should match number as a nested object head", func(t *testing.T) {
		pattern := st.NewNamedRulePattern("f", st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{
			st.NewAnonymousRulePattern(st.NewNumber("123"), st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{st.NewNumber("456")})),
		}))
		obj := base.NewNamedObject("f", []base.Node{base.NewObject(base.NewNumberFromString("123"), []base.Node{base.NewNumberFromString("456")})})

		assert.True(t, New(pattern)(obj).IsNotEmpty())
	})
}
