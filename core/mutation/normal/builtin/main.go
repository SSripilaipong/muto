package builtin

import (
	"github.com/SSripilaipong/muto/common/cliio"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

func NewMutators() []object.Mutator {
	return []object.Mutator{
		stringMutator,
		addMutator,
		subtractMutator,
		concatMutator,
		isStringMutator,
		isNumberMutator,
		isBooleanMutator,
		multiplyMutator,
		divideMutator,
		divIntegerMutator,
		modIntegerMutator,
		equalMutator,
		notEqualMutator,
		greaterThanMutator,
		greaterThanOrEqualMutator,
		lessThanMutator,
		lessThanOrEqualMutator,
		cliInputMutator(cliio.ReadInputOneLine),
		cliPrintMutator(cliio.PrintStringWithNewLine),
		andMutator,
		orMutator,
		notMutator,
		newTryMutator(),
	}
}
