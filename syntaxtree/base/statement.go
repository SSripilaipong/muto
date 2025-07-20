package base

type Statement interface {
	StatementType() StatementType
}

type StatementType string

const (
	RuleStatement       StatementType = "RULE"
	ActiveRuleStatement StatementType = "ACTIVE_RULE"
	ImportStatement     StatementType = "IMPORT"
)

func IsRuleStatement(s Statement) bool {
	return s.StatementType() == RuleStatement
}

func IsActiveRuleStatement(s Statement) bool {
	return s.StatementType() == ActiveRuleStatement
}
