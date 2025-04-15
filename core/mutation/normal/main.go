package normal

import (
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

func NewFromStatementsWithBuiltins(st []base.Statement, builtins []ruleMutator.NamedObjectMutator) ruleMutator.AppendableNameSwitch {
	return mergeNameWisedWithNamedObjectMutators(object.NewFromStatements(st), builtins)
}

func mergeNameWisedWithNamedObjectMutators(ms []ruleMutator.NameWrapper, ns []ruleMutator.NamedObjectMutator) ruleMutator.AppendableNameSwitch {
	var rs []ruleMutator.NamedObjectMutator
	for _, m := range ms {
		rs = append(rs, m)
	}
	return ruleMutator.NewAppendableNameSwitch(append(rs, ns...))
}
