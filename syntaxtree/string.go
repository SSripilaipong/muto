package syntaxtree

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

func (String) RuleResultType() RuleResultType {
	return RuleResultTypeString
}

func (s String) Value() string {
	return s.value
}
