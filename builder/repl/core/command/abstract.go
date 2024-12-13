package command

type Command interface {
	CommandType() Type
}

type Type string

const (
	TypeAddRule    Type = "ADD_RULE"
	TypeMutateNode Type = "MUTATE_NODE"
)

func IsAddRuleCommand(c Command) bool {
	return c.CommandType() == TypeAddRule
}

func IsMutateNodeCommand(c Command) bool {
	return c.CommandType() == TypeMutateNode
}

type TypeMixin struct {
	value Type
}

func NewTypeMixin(value Type) TypeMixin { return TypeMixin{value: value} }

func (mixin *TypeMixin) CommandType() Type { return mixin.value }
