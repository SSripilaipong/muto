package pattern

type FixedParamPart []Param

func (FixedParamPart) RulePatternParamPartType() ParamPartType {
	return ParamPartTypeFixed
}

func ParamsToFixedParamPart(xs []Param) FixedParamPart {
	return xs
}

func ParamsToParamPart(xs []Param) ParamPart {
	return ParamsToFixedParamPart(xs)
}

func UnsafeParamPartToParams(p ParamPart) []Param {
	return p.(FixedParamPart)
}
