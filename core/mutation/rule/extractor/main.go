package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

type topLevelNamedRule struct {
	param extractor.ParamChainPartial
}

func New(rule stPattern.NamedRule) mutator.Extractor {
	return topLevelNamedRule{param: newForParamChainPartial(stPattern.ExtractParamChain(rule))}
}

func (t topLevelNamedRule) Extract(x base.Object) optional.Of[*parameter.Parameter] {
	return t.param.Extract(x.ParamChain())
}

func newForParamChainPartial(chain []stPattern.ParamPart) extractor.ParamChainPartial {
	var extractors []extractor.NodeListExtractor
	if len(chain) > 0 {
		extractors = slc.Map(newForParamPartNested)(chain[:slc.LastIndex(chain)])
		extractors = append(extractors, newForParamPartTopLevel(slc.LastDefaultZero(chain)))
	}
	return extractor.NewParamChainPartial(extractors)
}

func newForParamPartTopLevel(paramPart stPattern.ParamPart) extractor.NodeListExtractor {
	switch {
	case stPattern.IsParamPartTypeFixed(paramPart):
		return newForFixedParamPartTopLevel(stPattern.UnsafeParamPartToParams(paramPart))
	case stPattern.IsParamPartTypeVariadic(paramPart):
		return newForVariadicParamPart(stPattern.UnsafeParamPartToVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func newForParamPartNested(paramPart stPattern.ParamPart) extractor.NodeListExtractor {
	switch {
	case stPattern.IsParamPartTypeFixed(paramPart):
		return newForFixedParamPart(stPattern.UnsafeParamPartToParams(paramPart))
	case stPattern.IsParamPartTypeVariadic(paramPart):
		return newForVariadicParamPart(stPattern.UnsafeParamPartToVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func newForVariadicParamPart(paramPart stPattern.VariadicParamPart) extractor.NodeListExtractor {
	switch {
	case stPattern.IsVariadicParamPartTypeRight(paramPart):
		return newForRightVariadicParamPart(stPattern.UnsafeVariadicParamPartToRightVariadicParamPart(paramPart))
	case stPattern.IsVariadicParamPartTypeLeft(paramPart):
		return newForLeftVariadicParamPart(stPattern.UnsafeVariadicParamPartToLeftVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func newForFixedParamPartTopLevel(params []stBase.PatternParam) extractor.ImplicitRightVariadic {
	return extractor.NewImplicitRightVariadic(newParamExtractors(params))
}

func newForFixedParamPart(params []stBase.PatternParam) extractor.NodeListExtractor {
	if len(params) == 0 {
		return extractor.NewExactNodeList(nil)
	}
	return extractor.NewExactNodeList(newParamExtractors(params))
}
