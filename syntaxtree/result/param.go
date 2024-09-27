package result

type Param interface {
	RuleResultNodeType() NodeType
	ObjectParamType() ParamType
}

func ParamToNode(x Param) Node {
	return x
}

type ParamType string

const (
	ParamTypeSingle   ParamType = "SINGLE"
	ParamTypeVariadic ParamType = "VARIADIC"
)

func IsParamTypeSingle(x Param) bool {
	return x.ObjectParamType() == ParamTypeSingle
}

func IsParamTypeVariadic(x Param) bool {
	return x.ObjectParamType() == ParamTypeVariadic
}
