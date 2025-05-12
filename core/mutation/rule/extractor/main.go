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

func New(rule stPattern.DeterminantObject) mutator.Extractor {
	return topLevelNamedRule{param: newForParamChainPartial(stPattern.ExtractParamChain(rule))}
}

func (t topLevelNamedRule) Extract(x base.Object) optional.Of[*parameter.Parameter] {
	return t.param.Extract(x.ParamChain())
}

func newForParamChainPartial(chain []stPattern.ParamPart) extractor.ParamChainPartial {
	var extractors []extractor.NodeListExtractor
	extractors = slc.Map(newForParamPartNested)(chain)
	if len(chain) > 0 {
		extractors = slc.Map(newForParamPartNested)(chain[:slc.LastIndex(chain)])
		if rightMostExtractor, ok := NewForParamPartTopLevel(slc.LastDefaultZero(chain)).Return(); ok {
			extractors = append(extractors, rightMostExtractor)
		}
	}
	return extractor.NewParamChainPartial(extractors)
}

func NewForParamPartTopLevel(paramPart stPattern.ParamPart) optional.Of[extractor.NodeListExtractor] {
	switch {
	case stPattern.IsParamPartTypeFixed(paramPart):
		return optional.Fmap(extractor.ToNodeListExtractor[extractor.ImplicitRightVariadic])(
			newForFixedParamPartTopLevel(stPattern.UnsafeParamPartToPatterns(paramPart)),
		)
	case stPattern.IsParamPartTypeVariadic(paramPart):
		return optional.Value(newForVariadicParamPart(stPattern.UnsafeParamPartToVariadicParamPart(paramPart)))
	}
	panic("not implemented")
}

func newForParamPartNested(paramPart stPattern.ParamPart) extractor.NodeListExtractor {
	switch {
	case stPattern.IsParamPartTypeFixed(paramPart):
		return newForFixedParamPart(stPattern.UnsafeParamPartToPatterns(paramPart))
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

func newForFixedParamPartTopLevel(params []stBase.Pattern) optional.Of[extractor.ImplicitRightVariadic] {
	if len(params) == 0 {
		return optional.Empty[extractor.ImplicitRightVariadic]()
	}
	return optional.Value(extractor.NewImplicitRightVariadic(newParamExtractors(params)))
}

func newForFixedParamPart(params []stBase.Pattern) extractor.NodeListExtractor {
	if len(params) == 0 {
		return extractor.NewExactNodeList(nil)
	}
	return extractor.NewExactNodeList(newParamExtractors(params))
}
