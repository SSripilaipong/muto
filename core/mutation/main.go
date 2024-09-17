package mutation

import (
	"muto/common/optional"
	"muto/core/base"
	activeMutation "muto/core/mutation/active"
	normalMutation "muto/core/mutation/normal"
	st "muto/syntaxtree"
)

func NewFromStatements(ss []st.Statement) func(base.MutableNode) optional.Of[base.Node] {
	mutate := mutation{
		active: activeMutation.NewFromStatements(ss),
		normal: normalMutation.NewFromStatements(ss),
	}
	return func(x base.MutableNode) optional.Of[base.Node] {
		return x.Mutate(mutate)
	}
}

type mutation struct {
	active func(name string, obj base.NamedObject) optional.Of[base.Node]
	normal func(name string, obj base.NamedObject) optional.Of[base.Node]
}

func (m mutation) Active(name string, obj base.NamedObject) optional.Of[base.Node] {
	return m.active(name, obj)
}

func (m mutation) Normal(name string, obj base.NamedObject) optional.Of[base.Node] {
	return m.normal(name, obj)
}
