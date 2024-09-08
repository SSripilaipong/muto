package base

type ParamPart interface{}

func NodesToParamPart(xs []Node) ParamPart {
	return xs
}
