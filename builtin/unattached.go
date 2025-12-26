package builtin

import (
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/portal"
)

type UnattachedImportMapping struct {
	mapping module.ImportMapping
}

func NewUnattachedImportMapping(mapping module.ImportMapping) UnattachedImportMapping {
	return UnattachedImportMapping{mapping: mapping}
}

func (m UnattachedImportMapping) Attach(q *portal.Portal) module.ImportMapping {
	m.mapping.MountPortal(q)
	return m.mapping
}
