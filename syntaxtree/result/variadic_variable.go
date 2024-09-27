package result

type VariadicVariable struct {
	name string
}

func NewVariadicVariable(name string) VariadicVariable {
	return VariadicVariable{name: name}
}

func (VariadicVariable) RuleResultNodeType() NodeType {
	return NodeTypeVariadicVariable
}

func (VariadicVariable) ObjectParamType() ParamType {
	return ParamTypeVariadic
}

func (v VariadicVariable) Name() string {
	return v.name
}

func UnsafeParamToVariadicVariable(x Param) VariadicVariable {
	return x.(VariadicVariable)
}
