package base

import "github.com/SSripilaipong/muto/common/optional"

type Class interface {
	Children() []Node
	NodeType() NodeType
	IsTerminationConfirmed() bool
	ConfirmTermination() MutableNode
	LiftTermination() MutableNode
	Mutate(mutation Mutation) optional.Of[Node]
	Name() string
}

var _ MutableNode = Class(nil)
