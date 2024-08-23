package syntaxtree

type Variable struct {
	name string
}

func NewVariable(name string) Variable {
	return Variable{name: name}
}

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

type Number struct {
	value string
}

func NewNumber(value string) RuleResult {
	return Number{value: value}
}
