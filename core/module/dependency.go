package module

import "github.com/SSripilaipong/muto/core/portal"

type Dependency struct {
	global   Base
	portal   *portal.Portal
	builtins ImportMapping
}

func NewDependency(global Base, q *portal.Portal, builtins ImportMapping) Dependency {
	return Dependency{
		global:   global,
		portal:   q,
		builtins: builtins,
	}
}

func (m Dependency) Global() Base {
	return m.global
}

func (m Dependency) Portal() *portal.Portal {
	return m.portal
}

func (m Dependency) Builtins() ImportMapping {
	return m.builtins
}
