package repl

import "github.com/SSripilaipong/muto/syntaxtree/base"

type ActiveRule struct {
	rule base.ActiveRule
}

func NewActiveRule(rule base.ActiveRule) ActiveRule {
	return ActiveRule{rule: rule}
}

func (r ActiveRule) ReplStatementType() StatementType {
	return StatementTypeActiveRule
}
