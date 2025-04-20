package extractor

import (
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newForRightVariadicParamPart(pp stPattern.RightVariadicParamPart) extractor.NodeListExtractor {
	if len(pp.Name()) == 0 {
		panic("right variadic param part's name should not be empty")
	}
	leftExtractors := newParamExtractors(pp.OtherPart())
	if pp.Name()[0] == '_' {
		return extractor.NewIgnoredRightVariadic(leftExtractors)
	}
	return extractor.NewRightVariadic(pp.Name(), leftExtractors)
}

func newForLeftVariadicParamPart(pp stPattern.LeftVariadicParamPart) extractor.NodeListExtractor {
	if len(pp.Name()) == 0 {
		panic("left variadic param part's name should not be empty")
	}
	rightExtractors := newParamExtractors(pp.OtherPart())
	if pp.Name()[0] == '_' {
		return extractor.NewIgnoredLeftVariadic(rightExtractors)
	}
	return extractor.NewLeftVariadic(pp.Name(), rightExtractors)
}
