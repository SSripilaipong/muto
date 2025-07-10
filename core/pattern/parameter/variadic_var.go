package parameter

import "github.com/SSripilaipong/muto/core/base"

type VariadicVarMapping struct {
	name  string
	nodes []base.Node
}

func NewVariadicVarMapping(name string, nodes []base.Node) VariadicVarMapping {
	return VariadicVarMapping{
		name:  name,
		nodes: nodes,
	}
}

func (x VariadicVarMapping) Name() string { return x.name }

func (x VariadicVarMapping) Nodes() []base.Node { return x.nodes }

func (x VariadicVarMapping) Equal(y VariadicVarMapping) bool {
	if y.name != x.name {
		return false
	}
	if len(x.nodes) != len(y.nodes) {
		return false
	}
	for i := range x.nodes {
		if !base.NodeEqual(x.nodes[i], y.nodes[i]) {
			return false
		}
	}
	return true
}
