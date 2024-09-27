package result

type FixedParamPart []Param

func (FixedParamPart) ObjectParamPartType() ParamPartType {
	return ParamPartTypeFixed
}

func (p FixedParamPart) Size() int {
	return len(p)
}

func UnsafeParamPartToFixedParamPart(part ParamPart) FixedParamPart {
	return part.(FixedParamPart)
}

func ParamsToFixedParamPart(params []Param) FixedParamPart {
	return params
}

func ParamsToParamPart(params []Param) ParamPart {
	return ParamsToFixedParamPart(params)
}
