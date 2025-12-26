package repl

import "strings"

type ImportCommand struct {
	commandTypeMixin
	statementTypeMixin
	path []string
}

func NewImportCommand(path []string) ImportCommand {
	return ImportCommand{
		statementTypeMixin: newStatementTypeMixin(StatementTypeReplCommand),
		commandTypeMixin:   newCommandTypeMixin(CommandTypeImport),
		path:               path,
	}
}

func (c ImportCommand) Path() []string {
	return c.path
}

func (c ImportCommand) JoinedPath() string {
	return strings.Join(c.path, ".")
}
