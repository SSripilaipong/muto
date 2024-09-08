package syntaxtree

type ObjectLeftVariadicParamPart struct {
	name      string
	otherPart ObjectFixedParamPart
}

func NewObjectLeftVariadicParamPart(name string, otherParam ObjectFixedParamPart) ObjectLeftVariadicParamPart {
	return ObjectLeftVariadicParamPart{name: name, otherPart: otherParam}
}

func (ObjectLeftVariadicParamPart) ObjectParamPartType() ObjectParamPartType {
	return ObjectParamPartTypeLeftVariadic
}

func (p ObjectLeftVariadicParamPart) Name() string {
	return p.name
}

func (p ObjectLeftVariadicParamPart) OtherPart() ObjectFixedParamPart {
	return p.otherPart
}

func UnsafeObjectParamPartToObjectLeftVariadicParamPart(part ObjectParamPart) ObjectLeftVariadicParamPart {
	return part.(ObjectLeftVariadicParamPart)
}

type ObjectRightVariadicParamPart struct {
	name      string
	otherPart ObjectFixedParamPart
}

func NewObjectRightVariadicParamPart(name string, otherParam ObjectFixedParamPart) ObjectRightVariadicParamPart {
	return ObjectRightVariadicParamPart{name: name, otherPart: otherParam}
}

func (ObjectRightVariadicParamPart) ObjectParamPartType() ObjectParamPartType {
	return ObjectParamPartTypeRightVariadic
}

func (p ObjectRightVariadicParamPart) Name() string {
	return p.name
}

func (p ObjectRightVariadicParamPart) OtherPart() ObjectFixedParamPart {
	return p.otherPart
}

func UnsafeObjectParamPartToObjectRightVariadicParamPart(part ObjectParamPart) ObjectRightVariadicParamPart {
	return part.(ObjectRightVariadicParamPart)
}
