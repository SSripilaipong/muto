package command

type ImportCommand struct {
	TypeMixin
	name string
}

func NewImport(name string) *ImportCommand {
	return &ImportCommand{
		TypeMixin: NewTypeMixin(TypeImport),
		name:      name,
	}
}

func (c ImportCommand) Name() string {
	return c.name
}

func UnsafeCommandToImportCommand(cmd Command) *ImportCommand {
	return cmd.(*ImportCommand)
}
