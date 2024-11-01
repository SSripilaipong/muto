package extractor

import (
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newForRightVariadicParamPart(pp stPattern.RightVariadicParamPart) extractor.RightVariadic {
	return extractor.NewRightVariadic(pp.Name(), newParamExtractors(pp.OtherPart()))
}

func newForLeftVariadicParamPart(pp stPattern.LeftVariadicParamPart) extractor.LeftVariadic {
	return extractor.NewLeftVariadic(pp.Name(), newParamExtractors(pp.OtherPart()))
}
