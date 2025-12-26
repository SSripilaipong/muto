package time

import (
	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

const ModuleName = "time"

func NewModule() module.Base {
	mod := global.NewBaseModule()
	normal := []mutator.NamedUnit{newSleepMutator()}

	collection := mutator.NewCollectionFromMutators(normal, nil)
	mod.ExtendCollection(collection)
	return mod
}
