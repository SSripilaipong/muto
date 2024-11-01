package parameter

import "github.com/SSripilaipong/muto/core/base"

type VariableMapping struct {
	name string
	node base.Node
}

func (m VariableMapping) Equals(n VariableMapping) bool {
	return m.name == n.name && base.NodeEqual(m.node, n.node)
}

func NewVariableMapping(name string, node base.Node) VariableMapping {
	return VariableMapping{
		name: name,
		node: node,
	}
}
