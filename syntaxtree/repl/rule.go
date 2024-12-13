package repl

import "github.com/SSripilaipong/muto/syntaxtree/base"

type Rule struct {
	rule base.Rule
}

func NewRule(rule base.Rule) Rule {
	return Rule{rule: rule}
}

func (r Rule) ReplStatementType() StatementType {
	return StatementTypeRule
}

func (r Rule) Rule() base.Rule {
	return r.rule
}

func UnsafeStatementToRule(s Statement) Rule {
	return s.(Rule)
}
