package extractor

import (
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newNestedNamedRuleExtractor(p stPattern.NamedRule) extractor.NodeExtractor {
	return extractor.NewObject(extractor.NewClass(p.ObjectName()), newForParamPartNested(p.ParamPart()))
}

func newNestedVariableRuleExtractor(p stPattern.VariableRule) extractor.NodeExtractor {
	return extractor.NewObject(extractor.NewVariable(p.VariableName()), newForParamPartNested(p.ParamPart()))
}

func newNestedAnonymousRuleExtractor(p stPattern.AnonymousRule) extractor.NodeExtractor {
	return extractor.NewObject(newParamExtractor(p.Head()), newForParamPartNested(p.ParamPart()))
}
