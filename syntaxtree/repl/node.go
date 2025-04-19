package repl

import (
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Node struct {
	statementTypeMixin
	node stResult.SimplifiedNode
}

func NewNode(node stResult.SimplifiedNode) Node {
	return Node{
		statementTypeMixin: newStatementTypeMixin(StatementTypeNode),
		node:               node,
	}
}

func (n Node) Node() stResult.SimplifiedNode {
	return n.node
}

func UnsafeStatementToNode(s Statement) Node {
	return s.(Node)
}
