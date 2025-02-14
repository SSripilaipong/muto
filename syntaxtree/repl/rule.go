package repl

import (
	"github.com/SSripilaipong/muto/syntaxtree"
)

type Rule struct {
	statementTypeMixin
	rule syntaxtree.Rule
}

func NewRule(rule syntaxtree.Rule) Rule {
	return Rule{
		statementTypeMixin: newStatementTypeMixin(StatementTypeRule),
		rule:               rule,
	}
}

func (r Rule) Rule() syntaxtree.Rule {
	return r.rule
}

func UnsafeStatementToRule(s Statement) Rule {
	return s.(Rule)
}
