package syntaxtree

type ObjectFixedParamPart []ObjectParam

func (ObjectFixedParamPart) ObjectParamPartType() ObjectParamPartType {
	return ObjectParamPartTypeFixed
}

func UnsafeObjectParamPartToObjectFixedParamPart(part ObjectParamPart) ObjectFixedParamPart {
	return part.(ObjectFixedParamPart)
}

func ObjectParamsToObjectFixedParamPart(params []ObjectParam) ObjectParamPart {
	return ObjectFixedParamPart(params)
}
