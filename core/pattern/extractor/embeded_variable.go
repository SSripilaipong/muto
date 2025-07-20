package extractor

import (
	"github.com/SSripilaipong/muto/common/rods"
	"github.com/SSripilaipong/muto/core/base"
)

type EmbeddedVariableFactory struct {
	variable         VariableFactory
	embeddedNode     rods.Map[string, base.Node]
	embeddedNodeList rods.Map[string, []base.Node]
}

func NewEmbeddedVariableFactory(embeddedNode rods.Map[string, base.Node], embeddedNodeList rods.Map[string, []base.Node]) EmbeddedVariableFactory {
	return EmbeddedVariableFactory{
		variable:         NewVariableFactory(),
		embeddedNode:     embeddedNode,
		embeddedNodeList: embeddedNodeList,
	}
}

func (m EmbeddedVariableFactory) FixedVariable(name string) NodeExtractor {
	if node, isEmbedded := m.embeddedNode.GetCollection(name).Return(); isEmbedded {
		return NewMatchSameNode(node)
	}
	return m.variable.FixedVariable(name)
}

func (m EmbeddedVariableFactory) VariadicVariable(name string) NodeListExtractor {
	if nodes, isEmbedded := m.embeddedNodeList.GetCollection(name).Return(); isEmbedded {
		return NewMatchSameNodeList(nodes)
	}
	return m.variable.VariadicVariable(name)
}
