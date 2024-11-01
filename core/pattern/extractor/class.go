package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Class struct {
	name string
}

func NewClass(name string) Class {
	return Class{name: name}
}

func (t Class) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.IsClassNode(node) && base.UnsafeNodeToClass(node).Name() == t.Name() {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (t Class) Name() string {
	return t.name
}

var _ NodeExtractor = Class{}
