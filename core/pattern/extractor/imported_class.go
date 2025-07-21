package extractor

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type ImportedClass struct {
	module string
	name   string
}

func NewImportedClass(module, name string) ImportedClass {
	return ImportedClass{module: module, name: name}
}

func (t ImportedClass) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.IsClassNode(node) {
		class := base.UnsafeNodeToClass(node)
		if base.IsImportedClass(class) {
			imported := base.UnsafeClassToImportedClass(class)
			if imported.Name() == t.Name() && imported.Module() == t.module {
				return optional.Value(parameter.New())
			}
		}
	}
	return optional.Empty[*parameter.Parameter]()
}

func (t ImportedClass) Name() string {
	return t.name
}

func (t ImportedClass) DisplayString() string {
	return t.Name()
}

var _ NodeExtractor = ImportedClass{}
