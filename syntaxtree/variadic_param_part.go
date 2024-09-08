package syntaxtree

type ObjectLeftVariadicParamPart struct {
	name       string
	otherParam ObjectParamPart
}

func NewObjectLeftVariadicParamPart(name string, otherParam ObjectParamPart) ObjectLeftVariadicParamPart {
	return ObjectLeftVariadicParamPart{name: name, otherParam: otherParam}
}

func (ObjectLeftVariadicParamPart) ObjectParamPartType() ObjectParamPartType {
	return ObjectParamPartTypeLeftVariadic
}

type ObjectRightVariadicParamPart struct {
	name       string
	otherParam ObjectParamPart
}

func NewObjectRightVariadicParamPart(name string, otherParam ObjectParamPart) ObjectRightVariadicParamPart {
	return ObjectRightVariadicParamPart{name: name, otherParam: otherParam}
}

func (ObjectRightVariadicParamPart) ObjectParamPartType() ObjectParamPartType {
	return ObjectParamPartTypeRightVariadic
}
