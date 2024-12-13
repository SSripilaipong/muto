package builtin

import (
	"github.com/SSripilaipong/muto/common/cliio"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

func NewMutators() []mutator.NamedObjectMutator {
	return []mutator.NamedObjectMutator{
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
		doMutator,
	}
}
