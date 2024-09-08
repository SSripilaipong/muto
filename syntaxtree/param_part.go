package syntaxtree

type ObjectParamPart interface {
	ObjectParamPartType() ObjectParamPartType
}

type ObjectParamPartType string

const (
	ObjectParamPartTypeFixed         ObjectParamPartType = "FIXED"
	ObjectParamPartTypeLeftVariadic  ObjectParamPartType = "LEFT_VARIADIC"
	ObjectParamPartTypeRightVariadic ObjectParamPartType = "RIGHT_VARIADIC"
)

func IsObjectParamPartTypeFixed(x ObjectParamPart) bool {
	return x.ObjectParamPartType() == ObjectParamPartTypeFixed
}

func IsObjectParamPartTypeLeftVariadic(x ObjectParamPart) bool {
	return x.ObjectParamPartType() == ObjectParamPartTypeLeftVariadic
}

func IsObjectParamPartTypeRightVariadic(x ObjectParamPart) bool {
	return x.ObjectParamPartType() == ObjectParamPartTypeRightVariadic
}
