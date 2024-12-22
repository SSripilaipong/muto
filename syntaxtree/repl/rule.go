package repl

import "github.com/SSripilaipong/muto/syntaxtree/base"

type Rule struct {
	statementTypeMixin
	rule base.Rule
}

func NewRule(rule base.Rule) Rule {
	return Rule{
		statementTypeMixin: newStatementTypeMixin(StatementTypeRule),
		rule:               rule,
	}
}

func (r Rule) Rule() base.Rule {
	return r.rule
}

func UnsafeStatementToRule(s Statement) Rule {
	return s.(Rule)
}
