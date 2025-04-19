package result

type ParamPart interface {
	ObjectParamPartType() ParamPartType // TODO remove this, just use ParamPartTypeFixed
}

type ParamPartType string

const (
	ParamPartTypeFixed ParamPartType = "FIXED"
)

func IsParamPartTypeFixed(x ParamPart) bool {
	return x.ObjectParamPartType() == ParamPartTypeFixed
}
