package repl

type Statement interface {
	ReplStatementType() StatementType
}

type StatementType string

const (
	StatementTypeReplCommand StatementType = "REPL_COMMAND"
	StatementTypeRule        StatementType = "RULE"
	StatementTypeActiveRule  StatementType = "ACTIVE_RULE"
	StatementTypeNode        StatementType = "NODE"
)

func ToStatement[T Statement](x T) Statement {
	return x
}

func IsReplCommand(s Statement) bool {
	return s.ReplStatementType() == StatementTypeReplCommand
}

func IsRuleStatement(s Statement) bool {
	return s.ReplStatementType() == StatementTypeRule
}

func IsActiveRuleStatement(s Statement) bool {
	return s.ReplStatementType() == StatementTypeActiveRule
}

func IsNodeStatement(s Statement) bool {
	return s.ReplStatementType() == StatementTypeNode
}

type statementTypeMixin struct {
	stmtType StatementType
}

func newStatementTypeMixin(stmtType StatementType) statementTypeMixin {
	return statementTypeMixin{stmtType: stmtType}
}

func (t statementTypeMixin) ReplStatementType() StatementType {
	return t.stmtType
}
