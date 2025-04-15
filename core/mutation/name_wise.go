package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	activeMutation "github.com/SSripilaipong/muto/core/mutation/active"
	normalMutation "github.com/SSripilaipong/muto/core/mutation/normal"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

type nameWiseMutation struct {
	active mutator.AppendableNameSwitch
	normal mutator.AppendableNameSwitch
}

func newNameWiseMutation(ss []st.Statement, builtins []mutator.NamedObjectMutator) nameWiseMutation {
	return nameWiseMutation{
		active: activeMutation.NewFromStatements(ss),
		normal: normalMutation.NewFromStatementsWithBuiltins(ss, builtins),
	}
}

func (m nameWiseMutation) Active(name string, obj base.Object) optional.Of[base.Node] {
	return m.active.MutateByName(name, obj)
}

func (m nameWiseMutation) Normal(name string, obj base.Object) optional.Of[base.Node] {
	return m.normal.MutateByName(name, obj)
}

func (m nameWiseMutation) AppendNormalRule(rule mutator.NamedObjectMutator) {
	m.normal.Append(rule)
}
