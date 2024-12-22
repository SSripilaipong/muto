package repl

type QuitCommand struct {
	commandTypeMixin
	statementTypeMixin
}

func NewQuitCommand() QuitCommand {
	return QuitCommand{
		statementTypeMixin: newStatementTypeMixin(StatementTypeReplCommand),
		commandTypeMixin:   newCommandTypeMixin(CommandTypeQuit),
	}
}
