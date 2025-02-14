package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type PrimitiveParam struct {
	extractor NodeExtractor
}

func NewPrimitiveParam(extractor NodeExtractor) PrimitiveParam {
	return PrimitiveParam{extractor: extractor}
}

func (p PrimitiveParam) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if !base.IsObjectNode(node) {
		return optional.Empty[*parameter.Parameter]()
	}
	obj := base.UnsafeNodeToObject(node)

	if obj.ParamChain().Size() != 1 {
		return optional.Empty[*parameter.Parameter]()
	}

	if len(obj.ParamChain().DirectParams()) != 0 {
		return optional.Empty[*parameter.Parameter]()
	}

	return p.extractor.Extract(obj.Head())
}
