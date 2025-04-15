package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

type TopLevelMutation struct {
	nameWiseMutation
}

func NewFromStatements(ss []st.Statement, builtins []mutator.NamedObjectMutator) TopLevelMutation {
	return TopLevelMutation{
		nameWiseMutation: newNameWiseMutation(ss, builtins),
	}
}

func (m TopLevelMutation) Mutate(x base.MutableNode) optional.Of[base.Node] {
	return x.Mutate(m.nameWiseMutation)
}
