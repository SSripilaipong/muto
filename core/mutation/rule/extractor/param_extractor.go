package extractor

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newParamExtractors(params []stPattern.Param) []extractor.NodeExtractor {
	return slc.Map(newParamExtractor)(params)
}

func newParamExtractor(p stPattern.Param) extractor.NodeExtractor {
	switch {
	case stPattern.IsParamTypeVariable(p):
		return newVariableParamExtractor(base.UnsafeRuleParamPatternToVariable(p))
	case stPattern.IsParamTypeBoolean(p):
		return newBooleanParamExtractor(base.UnsafeRuleParamPatternToBoolean(p))
	case stPattern.IsParamTypeString(p):
		return newStringParamExtractor(base.UnsafeRuleParamPatternToString(p))
	case stPattern.IsParamTypeNumber(p):
		return newNumberParamExtractor(base.UnsafeRuleParamPatternToNumber(p))
	case stPattern.IsParamTypeTag(p):
		return newTagParamExtractor(base.UnsafeRuleParamPatternToTag(p))
	case stPattern.IsParamTypeNestedNamedRule(p):
		return newNestedNamedRuleExtractor(stPattern.UnsafeParamToNamedRule(p))
	case stPattern.IsParamTypeNestedVariableRule(p):
		return newNestedVariableRuleExtractor(stPattern.UnsafeRuleParamPatternToVariableRulePattern(p))
	case stPattern.IsParamTypeNestedAnonymousRule(p):
		return newNestedAnonymousRuleExtractor(stPattern.UnsafeParamToAnonymousRule(p))
	}
	panic("not implemented")
}

func newVariableParamExtractor(v base.Variable) extractor.NodeExtractor {
	return extractor.NewVariable(v.Name())
}

func newBooleanParamExtractor(v base.Boolean) extractor.NodeExtractor {
	return extractor.NewBoolean(v.BooleanValue())
}

func newStringParamExtractor(v base.String) extractor.NodeExtractor {
	return extractor.NewString(v.StringValue())
}

func newNumberParamExtractor(v base.Number) extractor.NodeExtractor {
	return extractor.NewNumber(v.NumberValue())
}

func newTagParamExtractor(v base.Tag) extractor.NodeExtractor {
	return extractor.NewTag(v.Name())
}
