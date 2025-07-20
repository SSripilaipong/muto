package module

import (
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/portal"
)

type Module interface {
	GetClass(name string) base.Class
	LoadGlobal(global Module)
	MountPortal(q *portal.Portal)
	MutatorCollection() mutator.Collection
	MapImportedModules(mapping ImportMapping)
}
