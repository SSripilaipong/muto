package module

import (
	"github.com/SSripilaipong/muto/core/portal"
)

type Unattached[M Attachable] struct {
	module M
}

func NewUnattached[M Attachable](module M) Unattached[M] {
	return Unattached[M]{module: module}
}

func (m Unattached[M]) Attach(dep Dependency) M {
	m.module.MountPortal(dep.Portal())
	m.module.MapImportedModules(dep.Builtins())
	return m.module
}

type Attachable interface {
	MountPortal(q *portal.Portal)
	MapImportedModules(mapping ImportMapping)
}
