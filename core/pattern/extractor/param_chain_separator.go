package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type ParamChainSeparator struct {
	wrapped NodeListExtractor
}

func NewParamChainSeparator(wrapped NodeListExtractor) ParamChainSeparator {
	return ParamChainSeparator{wrapped: wrapped}
}

func (p ParamChainSeparator) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	return optional.Fmap(parameter.AddRemainingParamChain([]base.Node{}))(p.wrapped.Extract(nodes))
}

func (p ParamChainSeparator) DisplayString() string {
	return DisplayString(p.wrapped)
}

var _ NodeListExtractor = ParamChainSeparator{}
