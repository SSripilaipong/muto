package extractor

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

type namedRule struct {
	param extractor.ParamChainPartial
}

func NewNamedRule(rule stPattern.DeterminantObject) mutator.Extractor {
	variable := extractor.NewVariableFactory()
	core := newCorePatternFactory(variable)
	topLevel := NewTopLevelFactory(variable)
	return namedRule{param: newForParamChainPartial(core, topLevel, stPattern.ExtractParamChain(rule))}
}

func (r namedRule) Extract(x base.Object) optional.Of[*parameter.Parameter] {
	return r.param.Extract(x.ParamChain())
}

func newForParamChainPartial(core corePatternFactoryImpl, topLevel TopLevelFactory, chain []stPattern.ParamPart) extractor.ParamChainPartial {
	var extractors []extractor.NodeListExtractor
	if len(chain) > 0 {
		extractors = slc.Map(core.ParamPart)(chain[:slc.LastIndex(chain)])
		if rightMostExtractor, ok := topLevel.TopLevel(slc.LastDefaultZero(chain)).Return(); ok {
			extractors = append(extractors, rightMostExtractor)
		}
	}
	return extractor.NewParamChainPartial(extractors)
}
