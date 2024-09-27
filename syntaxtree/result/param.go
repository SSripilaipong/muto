package result

type Param interface {
	ObjectParamType() ParamType
}

func UnsafeParamToNode(x Param) Node {
	return x.(Node)
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
