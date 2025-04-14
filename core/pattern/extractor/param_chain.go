package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type ParamChainPartial struct {
	extractors []NodeListExtractor
}

func NewParamChainPartial(extractors []NodeListExtractor) ParamChainPartial {
	return ParamChainPartial{extractors}
}

func (p ParamChainPartial) Extract(x base.ParamChain) optional.Of[*parameter.Parameter] {
	merged := optional.Value(parameter.New())
	if len(p.extractors) > 0 {
		extractParams := slc.LeftZipApply(mapExtractNodeList(p.extractors))(optional.Empty[*parameter.Parameter]())
		params := extractParams(x.All())
		merged = slc.Reduce(optionalMergeParam)(params)
	}
	return p.chainRemainingParams(merged, x)
}

func (p ParamChainPartial) Size() int {
	return len(p.extractors)
}

func (p ParamChainPartial) chainRemainingParams(patternParam optional.Of[*parameter.Parameter], paramChain base.ParamChain) optional.Of[*parameter.Parameter] {
	pp, exists := patternParam.Return()
	if !exists || len(p.extractors) >= paramChain.Size() {
		return patternParam
	}
	remaining := paramChain.SliceFromOrEmpty(len(p.extractors)) // never out of range
	return optional.Value(pp.AppendAllRemainingParamChain(remaining))
}

var _ ParamChainExtractor = ParamChainPartial{}

type ParamChain struct {
	extractors []NodeListExtractor
}

func NewParamChain(extractors []NodeListExtractor) ParamChain {
	return ParamChain{extractors: extractors}
}

func (p ParamChain) Extract(x base.ParamChain) optional.Of[*parameter.Parameter] {
	if x.Size() != len(p.extractors) {
		return optional.Empty[*parameter.Parameter]()
	}
	extractParams := slc.LeftZipApply(mapExtractNodeList(p.extractors))(optional.Empty[*parameter.Parameter]())
	params := extractParams(x.All())
	return slc.Reduce(optionalMergeParam)(params)
}

func (p ParamChain) Size() int {
	return len(p.extractors)
}

var _ ParamChainExtractor = ParamChain{}

var mapExtractNodeList = slc.Map(extractNodeList)
