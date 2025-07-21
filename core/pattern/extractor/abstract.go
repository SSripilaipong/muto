package extractor

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type NodeExtractor interface {
	Extract(node base.Node) optional.Of[*parameter.Parameter]
}

func ToNodeExtractor[T NodeExtractor](extractor T) NodeExtractor {
	return extractor
}

type NodeListExtractor interface {
	Extract(nodes []base.Node) optional.Of[*parameter.Parameter]
}

func ToNodeListExtractor[T NodeListExtractor](extractor T) NodeListExtractor {
	return extractor
}

func extractNodeList(extractor NodeListExtractor) func(nodes []base.Node) optional.Of[*parameter.Parameter] {
	return extractor.Extract
}

type ParamChainExtractor interface {
	Extract(paramChain base.ParamChain) optional.Of[*parameter.Parameter]
	Size() int
	WrapDisplayString(head string) string
}
