package extractor

import (
	"fmt"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Object struct {
	headExtractor  NodeExtractor
	paramExtractor ParamChainExtractor
}

func NewObject(head NodeExtractor, param ParamChainExtractor) NodeExtractor {
	return Object{
		headExtractor:  head,
		paramExtractor: param,
	}
}

func (t Object) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	switch {
	case base.IsObjectNode(node):
		return t.extractObject(base.UnsafeNodeToObject(node))
	case base.IsClassNode(node):
		return t.extractObject(base.NewCompoundObject(node, base.NewParamChain([][]base.Node{})))
	}
	return optional.Empty[*parameter.Parameter]()
}

func (t Object) extractObject(obj base.Object) optional.Of[*parameter.Parameter] {
	paramChain := obj.ParamChain()
	if paramChain.Size() < t.paramExtractor.Size() {
		return optional.Empty[*parameter.Parameter]()
	}

	head := obj.Head()
	headSideParamChainSize := paramChain.Size() - t.paramExtractor.Size()
	if headSideParamChainSize > 0 {
		head = base.NewCompoundObject(obj.Head(), paramChain.SliceUntilOrEmpty(headSideParamChainSize))
	}

	splitParamChain := paramChain.SliceFromOrEmpty(paramChain.Size() - t.paramExtractor.Size())
	headParam := t.headExtractor.Extract(head)
	paramChainParam := t.paramExtractor.Extract(splitParamChain)
	return optionalMergeParam(paramChainParam, headParam)
}

func (t Object) DisplayString() string {
	headString := DisplayString(t.headExtractor)
	return fmt.Sprintf("(%s)", t.paramExtractor.WrapDisplayString(headString))
}

var _ NodeExtractor = Object{}

var optionalMergeParam = optional.MergeFn(parameter.Merge)
