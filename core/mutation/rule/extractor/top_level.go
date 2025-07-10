package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

type TopLevelFactory struct {
	core     corePatternFactory
	variable variableParamPartFactory
}

func NewTopLevelFactory(variable variableFactory) TopLevelFactory {
	core := newCorePatternFactory(variable)
	return TopLevelFactory{
		core:     core,
		variable: core.VariableParamPart(),
	}
}

func (f TopLevelFactory) TopLevel(paramPart stPattern.ParamPart) optional.Of[extractor.NodeListExtractor] {
	switch {
	case stPattern.IsParamPartTypeFixed(paramPart):
		return optional.Fmap(extractor.ToNodeListExtractor[extractor.ImplicitRightVariadic])(
			f.newForFixedParamPartTopLevel(stPattern.UnsafeParamPartToPatterns(paramPart)),
		)
	case stPattern.IsParamPartTypeVariadic(paramPart):
		return optional.Value(f.newForVariadicParamPartTopLevel(stPattern.UnsafeParamPartToVariadicParamPart(paramPart)))
	}
	panic("not implemented")
}

func (f TopLevelFactory) newForVariadicParamPartTopLevel(paramPart stPattern.VariadicParamPart) extractor.NodeListExtractor {
	return extractor.NewParamChainSeparator(f.variable.VariadicParamPart(paramPart))
}

func (f TopLevelFactory) newForFixedParamPartTopLevel(params []stBase.Pattern) optional.Of[extractor.ImplicitRightVariadic] {
	if len(params) == 0 {
		return optional.Empty[extractor.ImplicitRightVariadic]()
	}
	return optional.Value(extractor.NewImplicitRightVariadic(f.core.Patterns(params)))
}
