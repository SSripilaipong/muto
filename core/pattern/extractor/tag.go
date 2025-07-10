package extractor

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Tag struct {
	name string
}

func NewTag(name string) Tag {
	return Tag{name: name}
}

func (t Tag) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.IsTagNode(node) && base.UnsafeNodeToTag(node).Name() == t.Name() {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (t Tag) Name() string {
	return t.name
}

func (t Tag) DisplayString() string {
	return fmt.Sprintf(".%s", t.Name())
}

var _ NodeExtractor = Tag{}
