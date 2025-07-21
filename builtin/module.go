package builtin

import (
	"fmt"

	"github.com/SSripilaipong/go-common/rods"

	timeMod "github.com/SSripilaipong/muto/builtin/time"
	"github.com/SSripilaipong/muto/core/module"
)

func NewBuiltinImportMapping(names []string) UnattachedImportMapping {
	mapping := make(map[string]module.Module)
	for _, name := range names {
		switch name {
		case timeMod.ModuleName:
			mapping[timeMod.ModuleName] = timeMod.NewModule()
		default:
			panic(fmt.Sprintf("unknown module %#v", name))
		}
	}
	return NewUnattachedImportMapping(module.NewImportMapping(rods.NewMap(mapping)))
}
