package syntaxtree

type ObjectFixedParamPart []ObjectParam

func (ObjectFixedParamPart) ObjectParamPartType() ObjectParamPartType {
	return ObjectParamPartTypeFixed
}

func UnsafeObjectParamPartToObjectFixedParamPart(part ObjectParamPart) ObjectFixedParamPart {
	return part.(ObjectFixedParamPart)
}

func ObjectParamsToObjectFixedParamPart(params []ObjectParam) ObjectFixedParamPart {
	return params
}

func ObjectParamsToObjectParamPart(params []ObjectParam) ObjectParamPart {
	return ObjectParamsToObjectFixedParamPart(params)
}
