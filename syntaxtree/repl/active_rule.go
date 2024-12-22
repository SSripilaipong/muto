package repl

import "github.com/SSripilaipong/muto/syntaxtree/base"

type ActiveRule struct {
	statementTypeMixin
	rule base.ActiveRule
}

func NewActiveRule(rule base.ActiveRule) ActiveRule {
	return ActiveRule{
		statementTypeMixin: newStatementTypeMixin(StatementTypeActiveRule),
		rule:               rule,
	}
}
