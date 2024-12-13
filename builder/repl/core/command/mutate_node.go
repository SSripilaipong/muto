package command

import (
	"github.com/SSripilaipong/muto/core/base"
)

type MutateNodeCommand struct {
	TypeMixin
	node base.Node
}

func NewMutateNode(node base.Node) *MutateNodeCommand {
	return &MutateNodeCommand{
		TypeMixin: NewTypeMixin(TypeMutateNode),
		node:      node,
	}
}

func (c MutateNodeCommand) InitialNode() base.Node {
	return c.node
}

func UnsafeCommandToMutateNodeCommand(cmd Command) *MutateNodeCommand {
	return cmd.(*MutateNodeCommand)
}
