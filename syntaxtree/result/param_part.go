package result

type ParamPart interface {
	ObjectParamPartType() ParamPartType
}

type ParamPartType string

const (
	ParamPartTypeFixed ParamPartType = "FIXED"
)

func IsParamPartTypeFixed(x ParamPart) bool {
	return x.ObjectParamPartType() == ParamPartTypeFixed
}
