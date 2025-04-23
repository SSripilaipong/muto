package result

import "slices"

type FixedParamPart []Param

func (p FixedParamPart) Size() int {
	return len(p)
}

func (p FixedParamPart) Append(q FixedParamPart) FixedParamPart {
	return append(slices.Clone(p), q...)
}

func ParamsToFixedParamPart(params []Param) FixedParamPart {
	return params
}
