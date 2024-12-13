package program

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

func (w Wrapper) AddRule(rule mutator.NamedObjectMutator) optional.Of[int] {
	w.program.AddRule(rule)
	return optional.Empty[int]()
}
