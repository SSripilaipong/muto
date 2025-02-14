package extractor

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newNestedNamedRuleExtractor(p stPattern.NamedRule) extractor.NodeExtractor {
	return extractor.NewObject(extractor.NewClass(p.ObjectName()), newForParamChain(p.ParamParts()))
}

func newNestedVariableRuleExtractor(p stPattern.VariableRule) extractor.NodeExtractor {
	return extractor.NewObject(extractor.NewHeadVariable(p.VariableName()), newForParamChain(p.ParamParts()))
}

func newNestedAnonymousRuleExtractor(p stPattern.AnonymousRule) extractor.NodeExtractor {
	return extractor.NewObject(newHeadExtractor(p.Head()), newForParamChain(p.ParamParts()))
}

func newForParamChain(chain []stPattern.ParamPart) extractor.ParamChain {
	extractors := slc.Map(newForParamPartNested)(chain)
	return extractor.NewParamChain(extractors)
}

func newHeadExtractor(p base.PatternParam) extractor.NodeExtractor {
	x, ok := newPrimitiveExtractor(p).Return()
	if ok {
		return x
	}
	return newNonPrimitiveExtractor(p)
}
