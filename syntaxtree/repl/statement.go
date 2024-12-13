package repl

type Statement interface {
	ReplStatementType() StatementType
}

type StatementType string

const (
	StatementTypeRule       StatementType = "RULE"
	StatementTypeActiveRule StatementType = "ACTIVE_RULE"
	StatementTypeNode       StatementType = "NODE"
)

func IsRuleStatement(s Statement) bool {
	return s.ReplStatementType() == StatementTypeRule
}

func IsActiveRuleStatement(s Statement) bool {
	return s.ReplStatementType() == StatementTypeActiveRule
}

func IsNodeStatement(s Statement) bool {
	return s.ReplStatementType() == StatementTypeNode
}
