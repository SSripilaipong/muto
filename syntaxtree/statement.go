package syntaxtree

type Statement interface {
	StatementType() StatementType
}

type StatementType string

const RuleStatement StatementType = "RULE"

func IsRuleStatement(s Statement) bool {
	return s.StatementType() == RuleStatement
}
