package active

import (
	"github.com/SSripilaipong/muto/common/fn"
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

var NewFromStatements = fn.Compose(mergeNameWised, newMutatorsFromStatements)

func mergeNameWised(ms []ruleMutator.NameWrapper) ruleMutator.AppendableNameSwitch {
	var rs []ruleMutator.NamedObjectMutator
	for _, m := range ms {
		rs = append(rs, m)
	}
	return ruleMutator.NewAppendableNameSwitch(rs)
}
