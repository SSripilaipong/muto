package pattern

type ParamPart interface {
	RulePatternParamPartType() ParamPartType
}

type ParamPartType string

const (
	ParamPartTypeFixed    ParamPartType = "FIXED"
	ParamPartTypeVariadic ParamPartType = "VARIADIC"
)

func IsParamPartTypeFixed(p ParamPart) bool {
	return p.RulePatternParamPartType() == ParamPartTypeFixed
}

func IsParamPartTypeVariadic(p ParamPart) bool {
	return p.RulePatternParamPartType() == ParamPartTypeVariadic
}
