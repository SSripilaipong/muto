package repl

import (
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Node struct {
	node stResult.Node
}

func NewNode(node stResult.Node) Node {
	return Node{node: node}
}

func (n Node) ReplStatementType() StatementType {
	return StatementTypeNode
}

func (n Node) Node() stResult.Node {
	return n.node
}

func UnsafeStatementToNode(s Statement) Node {
	return s.(Node)
}
