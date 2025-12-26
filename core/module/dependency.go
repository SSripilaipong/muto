package module

import "github.com/SSripilaipong/muto/core/portal"

type Dependency struct {
	portal   *portal.Portal
	builtins ImportMapping
}

func NewDependency(q *portal.Portal, builtins ImportMapping) Dependency {
	return Dependency{
		portal:   q,
		builtins: builtins,
	}
}

func (m Dependency) Portal() *portal.Portal {
	return m.portal
}

func (m Dependency) Builtins() ImportMapping {
	return m.builtins
}
