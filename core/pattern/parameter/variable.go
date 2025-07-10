package parameter

import "github.com/SSripilaipong/muto/core/base"

type VariableMapping struct {
	name string
	node base.Node
}

func (m VariableMapping) Equals(n VariableMapping) bool {
	return m.name == n.name && base.NodeEqual(m.node, n.node)
}

func (m VariableMapping) Name() string { return m.name }

func (m VariableMapping) Node() base.Node { return m.node }

func NewVariableMapping(name string, node base.Node) VariableMapping {
	return VariableMapping{
		name: name,
		node: node,
	}
}
