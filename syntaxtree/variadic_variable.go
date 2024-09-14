package syntaxtree

type VariadicVariable struct {
	name string
}

func NewVariadicVariable(name string) VariadicVariable {
	return VariadicVariable{name: name}
}

func (VariadicVariable) RuleResultType() RuleResultType {
	return RuleResultTypeVariadicVariable
}

func (VariadicVariable) ObjectParamType() ObjectParamType { return ObjectParamTypeVariadic }

func (v VariadicVariable) Name() string {
	return v.name
}

func UnsafeObjectParamToVariadicVariable(x ObjectParam) VariadicVariable {
	return x.(VariadicVariable)
}
