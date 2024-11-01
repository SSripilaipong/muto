package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type LeafObject struct {
	head NodeExtractor
}

func NewLeafObject(head NodeExtractor) LeafObject {
	return LeafObject{head: head}
}

func (t LeafObject) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if !base.IsObjectNode(node) {
		return optional.Empty[*parameter.Parameter]()
	}
	obj := base.UnsafeNodeToObject(node)
	if len(obj.Children()) > 0 {
		return optional.Empty[*parameter.Parameter]()
	}
	return t.head.Extract(obj.Head())
}

var _ NodeExtractor = LeafObject{}
