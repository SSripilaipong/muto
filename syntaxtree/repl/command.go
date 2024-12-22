package repl

type Command interface {
	CommandType() CommandType
}

type CommandType string

const (
	CommandTypeQuit CommandType = "QUIT"
)

func IsQuitCommand(cmd Command) bool {
	return cmd.CommandType() == CommandTypeQuit
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
