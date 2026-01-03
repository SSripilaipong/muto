package extractor

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

type corePatternFactory interface {
	Patterns(patterns []stBase.Pattern) []extractor.NodeExtractor
	ParamPart(paramPart stPattern.ParamPart) extractor.NodeListExtractor
}

type corePatternFactoryImpl struct {
	variable  variableParamPartFactory
	nonObject nonObjectFactory
}

func newCorePatternFactory(variable VariableFactory) corePatternFactoryImpl {
	core := corePatternFactoryImpl{}
	core.variable = newVariableParamPartFactory(&core, variable)
	core.nonObject = newNonObjectFactory(core.variable)
	return core
}

func (c corePatternFactoryImpl) Patterns(patterns []stBase.Pattern) []extractor.NodeExtractor {
	return slc.Map(c.newCorePatternExtractor)(patterns)
}

func (c corePatternFactoryImpl) ParamPart(paramPart stPattern.ParamPart) extractor.NodeListExtractor {
	switch {
	case stPattern.IsParamPartTypeFixed(paramPart):
		return c.newForFixedParamPart(stPattern.UnsafeParamPartToPatterns(paramPart))
	case stPattern.IsParamPartTypeVariadic(paramPart):
		return c.variable.VariadicParamPart(stPattern.UnsafeParamPartToVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func (c corePatternFactoryImpl) VariableParamPart() variableParamPartFactory {
	return c.variable
}

func (c corePatternFactoryImpl) newForFixedParamPart(params []stBase.Pattern) extractor.NodeListExtractor {
	if len(params) == 0 {
		return extractor.NewExactNodeList(nil)
	}
	return extractor.NewExactNodeList(c.Patterns(params))
}

func (c corePatternFactoryImpl) newCorePatternExtractor(p stBase.Pattern) extractor.NodeExtractor {
	if stBase.IsPatternTypeConjunction(p) {
		conj := stPattern.UnsafePatternToConjunction(p)
		return extractor.NewConjunction(
			c.newCorePatternExtractor(conj.Main()),
			c.newCorePatternExtractor(conj.Conj()),
		)
	}
	if r, ok := c.nonObject.TryNonObject(p).Return(); ok {
		return r
	}
	if stBase.IsPatternTypeObject(p) {
		return c.newObjectExtractor(stPattern.UnsafePatternToObject(p))
	}
	panic("not implemented")
}

func (c corePatternFactoryImpl) newObjectExtractor(p stPattern.Object) extractor.NodeExtractor {
	baseExtractor := extractor.NewObject(
		c.nonObject.NonObject(stPattern.ExtractNonObjectHead(p)),
		c.newForParamChain(stPattern.ExtractParamChain(p)),
	)

	// Extract head conjunctions and wrap with NLayeredConjunction if present
	headConjs := stPattern.ExtractHeadConjunctions(p)
	conjLayers := c.buildConjunctionLayers(headConjs)
	return extractor.NewNLayeredConjunction(baseExtractor, conjLayers)
}

func (c corePatternFactoryImpl) buildConjunctionLayers(headConjs [][]stBase.Pattern) [][]extractor.NodeExtractor {
	layers := make([][]extractor.NodeExtractor, len(headConjs))
	for i, conjs := range headConjs {
		if len(conjs) == 0 {
			layers[i] = nil
			continue
		}
		layers[i] = make([]extractor.NodeExtractor, len(conjs))
		for j, conj := range conjs {
			layers[i][j] = c.newCorePatternExtractor(conj)
		}
	}
	return layers
}

func (c corePatternFactoryImpl) newForParamChain(chain []stPattern.ParamPart) extractor.ParamChain {
	extractors := slc.Map(c.ParamPart)(chain)
	return extractor.NewParamChain(extractors)
}
