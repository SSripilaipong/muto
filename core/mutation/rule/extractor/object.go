package extractor

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newObjectExtractor(p stPattern.Object) extractor.NodeExtractor {
	return extractor.NewObject(
		newNonObjectPatternExtractor(stPattern.ExtractNonObjectHead(p)),
		newForParamChain(stPattern.ExtractParamChain(p)),
	)
}

func newForParamChain(chain []stPattern.ParamPart) extractor.ParamChain {
	extractors := slc.Map(newForParamPartNested)(chain)
	return extractor.NewParamChain(extractors)
}
