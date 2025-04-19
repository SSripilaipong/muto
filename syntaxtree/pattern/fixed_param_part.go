package pattern

import "github.com/SSripilaipong/muto/syntaxtree/base"

type FixedParamPart []base.Pattern

func (FixedParamPart) RulePatternParamPartType() ParamPartType {
	return ParamPartTypeFixed
}

func (p FixedParamPart) Size() int { return len(p) }

func PatternsToFixedParamPart(xs []base.Pattern) FixedParamPart {
	return xs
}

func PatternsToParamPart(xs []base.Pattern) ParamPart {
	return PatternsToFixedParamPart(xs)
}

func UnsafeParamPartToFixedParamPart(p ParamPart) FixedParamPart {
	return p.(FixedParamPart)
}

func UnsafeParamPartToPatterns(p ParamPart) []base.Pattern {
	return p.(FixedParamPart)
}
