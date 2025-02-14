package pattern

import "github.com/SSripilaipong/muto/syntaxtree/base"

type FixedParamPart []base.PatternParam

func (FixedParamPart) RulePatternParamPartType() ParamPartType {
	return ParamPartTypeFixed
}

func ParamsToFixedParamPart(xs []base.PatternParam) FixedParamPart {
	return xs
}

func ParamsToParamPart(xs []base.PatternParam) ParamPart {
	return ParamsToFixedParamPart(xs)
}

func UnsafeParamPartToParams(p ParamPart) []base.PatternParam {
	return p.(FixedParamPart)
}
