package builtin

import "phi-lang/core/mutation/object"

func NewMutators() []object.Mutator {
	return []object.Mutator{
		addMutator,
		concatMutator,
		stringMutator,
	}
}
