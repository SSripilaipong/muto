package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type LocalClass struct {
	name string
}

func NewLocalClass(name string) LocalClass {
	return LocalClass{name: name}
}

func (t LocalClass) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.IsClassNode(node) {
		class := base.UnsafeNodeToClass(node)
		if base.IsRuleBasedClass(class) && base.UnsafeClassToRuleBasedClass(class).Name() == t.Name() {
			return optional.Value(parameter.New())
		}
	}
	return optional.Empty[*parameter.Parameter]()
}

func (t LocalClass) Name() string {
	return t.name
}

func (t LocalClass) DisplayString() string {
	return t.Name()
}

var _ NodeExtractor = LocalClass{}
