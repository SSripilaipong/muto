package syntaxtree

type Number struct {
	value string
}

func (Number) RuleResultType() RuleResultType {
	return RuleResultTypeNumber
}

func (n Number) Value() string {
	return n.value
}

func NewNumber(value string) Number {
	return Number{value: value}
}
