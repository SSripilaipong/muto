package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newParamExtractors(params []base.PatternParam) []extractor.NodeExtractor {
	return slc.Map(newParamExtractor)(params)
}

func newParamExtractor(p base.PatternParam) extractor.NodeExtractor {
	x, ok := newPrimitiveExtractor(p).Return()
	if ok {
		return x
	}
	return newNonPrimitiveExtractor(p)
}

func newNonPrimitiveExtractor(p base.PatternParam) extractor.NodeExtractor {
	switch {
	case base.IsPatternParamTypeVariable(p):
		return newVariableParamExtractor(base.UnsafeRuleParamPatternToVariable(p))
	case base.IsPatternParamTypeNestedNamedRule(p):
		return newNestedNamedRuleExtractor(stPattern.UnsafeParamToNamedRule(p))
	case base.IsPatternParamTypeNestedVariableRule(p):
		return newNestedVariableRuleExtractor(stPattern.UnsafeRuleParamPatternToVariableRulePattern(p))
	case base.IsPatternParamTypeNestedAnonymousRule(p):
		return newNestedAnonymousRuleExtractor(stPattern.UnsafeParamToAnonymousRule(p))
	}
	panic("not implemented")
}

func newPrimitiveExtractor(p base.PatternParam) optional.Of[extractor.NodeExtractor] {
	switch {
	case base.IsPatternParamTypeBoolean(p):
		return optional.Value(newBooleanParamExtractor(base.UnsafeRuleParamPatternToBoolean(p)))
	case base.IsPatternParamTypeString(p):
		return optional.Value(newStringParamExtractor(base.UnsafeRuleParamPatternToString(p)))
	case base.IsPatternParamTypeNumber(p):
		return optional.Value(newNumberParamExtractor(base.UnsafeRuleParamPatternToNumber(p)))
	case base.IsPatternParamTypeTag(p):
		return optional.Value(newTagParamExtractor(base.UnsafeRuleParamPatternToTag(p)))
	}
	return optional.Empty[extractor.NodeExtractor]()
}

func newVariableParamExtractor(v base.Variable) extractor.NodeExtractor {
	return extractor.NewParamVariable(v.Name())
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
