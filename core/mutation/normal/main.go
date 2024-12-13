package normal

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/core/mutation/normal/builtin"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

var NewFromStatements = fn.Compose(mergeNameWised, object.NewMutatorsFromStatements)

func mergeNameWised(ms []ruleMutator.NameWrapper) ruleMutator.AppendableNameSwitch {
	var rs []ruleMutator.NamedObjectMutator
	for _, m := range ms {
		rs = append(rs, m)
	}
	return ruleMutator.NewAppendableNameSwitch(append(rs, builtin.NewMutators()...))
}
