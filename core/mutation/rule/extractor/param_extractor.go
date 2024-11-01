package extractor

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newParamExtractors(params []stPattern.Param) []extractor.NodeExtractor {
	return slc.Map(newParamExtractor)(params)
}

func newParamExtractor(p stPattern.Param) extractor.NodeExtractor {
	switch {
	case stPattern.IsParamTypeVariable(p):
		return newVariableParamExtractor(st.UnsafeRuleParamPatternToVariable(p))
	case stPattern.IsParamTypeBoolean(p):
		return newBooleanParamExtractor(st.UnsafeRuleParamPatternToBoolean(p))
	case stPattern.IsParamTypeString(p):
		return newStringParamExtractor(st.UnsafeRuleParamPatternToString(p))
	case stPattern.IsParamTypeNumber(p):
		return newNumberParamExtractor(st.UnsafeRuleParamPatternToNumber(p))
	case stPattern.IsParamTypeTag(p):
		return newTagParamExtractor(st.UnsafeRuleParamPatternToTag(p))
	case stPattern.IsParamTypeNestedNamedRule(p):
		return newNestedNamedRuleExtractor(stPattern.UnsafeParamToNamedRule(p))
	case stPattern.IsParamTypeNestedVariableRule(p):
		return newNestedVariableRuleExtractor(stPattern.UnsafeRuleParamPatternToVariableRulePattern(p))
	case stPattern.IsParamTypeNestedAnonymousRule(p):
		return newNestedAnonymousRuleExtractor(stPattern.UnsafeParamToAnonymousRule(p))
	}
	panic("not implemented")
}

func newVariableParamExtractor(v st.Variable) extractor.NodeExtractor {
	return extractor.NewVariable(v.Name())
}

func newBooleanParamExtractor(v st.Boolean) extractor.NodeExtractor {
	return extractor.NewBoolean(v.BooleanValue())
}

func newStringParamExtractor(v st.String) extractor.NodeExtractor {
	return extractor.NewString(v.StringValue())
}

func newNumberParamExtractor(v st.Number) extractor.NodeExtractor {
	return extractor.NewNumber(v.NumberValue())
}

func newTagParamExtractor(v st.Tag) extractor.NodeExtractor {
	return extractor.NewTag(v.Name())
}
