package syntaxtree

type ObjectParamPart interface {
	ObjectParamPartType() ObjectParamPartType
}

type ObjectParamPartType string

const (
	ObjectParamPartTypeFixed ObjectParamPartType = "FIXED"
)

func IsObjectParamPartTypeFixed(x ObjectParamPart) bool {
	return x.ObjectParamPartType() == ObjectParamPartTypeFixed
}
