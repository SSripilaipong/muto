package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type ReducibleObject struct {
	head     NodeExtractor
	children NodeListExtractor
}

func NewReducibleObject(head NodeExtractor, children NodeListExtractor) NodeExtractor {
	return ReducibleObject{
		head:     head,
		children: children,
	}
}

func (t ReducibleObject) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	switch {
	case base.IsObjectNode(node):
		return t.extractObject(base.UnsafeNodeToObject(node))
	case base.IsClassNode(node):
		return t.extractObject(base.NewObject(node, nil))
	}
	return optional.Empty[*parameter.Parameter]()
}

func (t ReducibleObject) extractObject(obj base.Object) optional.Of[*parameter.Parameter] {
	return optional.MergeFn(parameter.Merge)(
		t.children.Extract(obj.Children()), t.head.Extract(obj.Head()),
	)
}

var _ NodeExtractor = ReducibleObject{}
