package syntaxtree

type RuleResult interface {
	RuleResultType() RuleResultType
}

type RuleResultType string

const (
	RuleResultTypeString           RuleResultType = "STRING"
	RuleResultTypeNumber           RuleResultType = "NUMBER"
	RuleResultTypeNamedObject      RuleResultType = "NAMED_OBJECT"
	RuleResultTypeAnonymousObject  RuleResultType = "ANONYMOUS_OBJECT"
	RuleResultTypeVariable         RuleResultType = "VARIABLE"
	RuleResultTypeVariadicVariable RuleResultType = "VARIADIC_VARIABLE"
)

func IsRuleResultTypeString(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeString
}

func IsRuleResultTypeNumber(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeNumber
}

func IsRuleResultTypeAnonymousObject(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeAnonymousObject
}

func IsRuleResultTypeNamedObject(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeNamedObject
}

func IsRuleResultTypeVariable(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeVariable
}

type ObjectParam interface {
	RuleResultType() RuleResultType
	ObjectParamType() ObjectParamType
}

func ObjectParamToRuleResult(x ObjectParam) RuleResult {
	return x
}

type ObjectParamType string

const (
	ObjectParamTypeSingle   ObjectParamType = "SINGLE"
	ObjectParamTypeVariadic ObjectParamType = "VARIADIC"
)

func IsObjectParamTypeSingle(x ObjectParam) bool {
	return x.ObjectParamType() == ObjectParamTypeSingle
}

func IsObjectParamTypeVariadic(x ObjectParam) bool {
	return x.ObjectParamType() == ObjectParamTypeVariadic
}
