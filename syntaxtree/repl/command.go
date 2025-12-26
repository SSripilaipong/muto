package repl

type Command interface {
	CommandType() CommandType
}

type CommandType string

const (
	CommandTypeQuit   CommandType = "QUIT"
	CommandTypeImport CommandType = "IMPORT"
)

func IsQuitCommand(cmd Command) bool {
	return cmd.CommandType() == CommandTypeQuit
}

func IsImportCommand(cmd Command) bool {
	return cmd.CommandType() == CommandTypeImport
}

type commandTypeMixin struct {
	cmdType CommandType
}

func newCommandTypeMixin(cmdType CommandType) commandTypeMixin {
	return commandTypeMixin{cmdType: cmdType}
}

func (t commandTypeMixin) CommandType() CommandType {
	return t.cmdType
}

func UnsafeStatementToReplCommand(s Statement) Command {
	return s.(Command)
}

func UnsafeCommandToImportCommand(s Command) ImportCommand {
	return s.(ImportCommand)
}
