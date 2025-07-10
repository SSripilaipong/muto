package extractor

import (
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type MatchSameNode struct {
	node base.Node
}

func NewMatchSameNode(node base.Node) MatchSameNode {
	return MatchSameNode{node: node}
}

func (m MatchSameNode) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	if base.NodeEqual(m.node, node) {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (m MatchSameNode) DisplayString() string {
	return m.node.TopLevelString()
}

type MatchSameNodeList struct {
	nodes []base.Node
}

func NewMatchSameNodeList(nodes []base.Node) MatchSameNodeList {
	return MatchSameNodeList{nodes: nodes}
}

func (m MatchSameNodeList) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	if areTheSame := func() bool {
		if len(nodes) != len(m.nodes) {
			return false
		}
		for i, node := range nodes {
			if base.NodeNotEqual(m.nodes[i], node) {
				return false
			}
		}
		return true
	}(); areTheSame {
		return optional.Value(parameter.New())
	}
	return optional.Empty[*parameter.Parameter]()
}

func (m MatchSameNodeList) DisplayString() string {
	var result []string
	for _, node := range m.nodes {
		result = append(result, node.TopLevelString())
	}
	return strings.Join(result, " ")
}
