package module

import "github.com/SSripilaipong/muto/core/portal"

type Uninitialized struct {
	module Module
}

func AsUninitialized(module Module) Uninitialized {
	return Uninitialized{module: module}
}

func (m Uninitialized) Init(builtin Module, q *portal.Portal) Module {
	m.module.LoadGlobal(builtin)
	m.module.MountPortal(q)
	return m.module
}
