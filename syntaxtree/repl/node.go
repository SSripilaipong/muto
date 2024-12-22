package repl

import (
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Node struct {
	statementTypeMixin
	node stResult.Node
}

func NewNode(node stResult.Node) Node {
	return Node{
		statementTypeMixin: newStatementTypeMixin(StatementTypeNode),
		node:               node,
	}
}

func (n Node) Node() stResult.Node {
	return n.node
}

func UnsafeStatementToNode(s Statement) Node {
	return s.(Node)
}
