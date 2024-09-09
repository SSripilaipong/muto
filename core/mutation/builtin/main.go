package builtin

import "muto/core/mutation/object"

func NewMutators() []object.Mutator {
	return []object.Mutator{
		rollingAddMutator,
		rollingConcatMutator,
		stringMutator,
	}
}
