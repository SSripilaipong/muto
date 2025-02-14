package repl

import (
	"github.com/SSripilaipong/muto/syntaxtree"
)

type ActiveRule struct {
	statementTypeMixin
	rule syntaxtree.ActiveRule
}

func NewActiveRule(rule syntaxtree.ActiveRule) ActiveRule {
	return ActiveRule{
		statementTypeMixin: newStatementTypeMixin(StatementTypeActiveRule),
		rule:               rule,
	}
}
