package time

import (
	"github.com/SSripilaipong/muto/core/module"
	mutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

const ModuleName = "time"

func NewModule() module.Base {
	collection := mutator.NewCollectionFromMutators([]mutator.NamedUnit{newSleepMutator()}, nil)
	return module.NewBase(collection, mutation.NewRuleBuilder())
}
