package extractor

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

type extractorFunc func(base.Object) optional.Of[*parameter.Parameter]

func (f extractorFunc) Extract(x base.Object) optional.Of[*parameter.Parameter] {
	return f(x)
}

func New(rule stPattern.NamedRule) mutator.Extractor {
	return extractorFunc(fn.Compose(newForParamPartTopLevel(rule.ParamPart()).Extract, base.ObjectToChildren))
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

func newForFixedParamPartTopLevel(params []stPattern.Param) extractor.ImplicitRightVariadic {
	return extractor.NewImplicitRightVariadic(newParamExtractors(params))
}

func newForFixedParamPart(params []stPattern.Param) extractor.NodeListExtractor {
	if len(params) == 0 {
		return nil
	}
	return extractor.NewExactNodeList(newParamExtractors(params))
}
