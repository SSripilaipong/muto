package syntaxtree

type NamedRulePattern struct {
	objectName string
	params     RulePatternParamPart
}

func (NamedRulePattern) RuleParamPatternType() RuleParamPatternType {
	return RuleParamPatternTypeNestedNamedRulePattern
}

func (p NamedRulePattern) ObjectName() string {
	return p.objectName
}

func (p NamedRulePattern) ParamPart() RulePatternParamPart {
	return p.params
}

func NewNamedRulePattern(objectName string, params RulePatternParamPart) NamedRulePattern {
	return NamedRulePattern{objectName: objectName, params: params}
}

func NamedRulePatternToRulePatternParam(x NamedRulePattern) RuleParamPattern {
	return x
}

func UnsafeRuleParamPatternToNamedRulePattern(p RuleParamPattern) NamedRulePattern {
	return p.(NamedRulePattern)
}

type VariableRulePattern struct {
	variableName string
	params       RulePatternParamPart
}

func (VariableRulePattern) RuleParamPatternType() RuleParamPatternType {
	return RuleParamPatternTypeNestedVariableRulePattern
}

func (p VariableRulePattern) VariableName() string {
	return p.variableName
}

func (p VariableRulePattern) ParamPart() RulePatternParamPart {
	return p.params
}

func NewVariableRulePattern(variableName string, params RulePatternParamPart) VariableRulePattern {
	return VariableRulePattern{variableName: variableName, params: params}
}

func VariableRulePatternToRulePatternParam(x VariableRulePattern) RuleParamPattern {
	return x
}

type RulePatternParamPart interface {
	RulePatternParamPartType() RulePatternParamPartType
}

type RulePatternParamPartType string

const (
	RulePatternParamPartTypeFixed    RulePatternParamPartType = "FIXED"
	RulePatternParamPartTypeVariadic RulePatternParamPartType = "VARIADIC"
)

func IsRulePatternParamPartTypeFixed(p RulePatternParamPart) bool {
	return p.RulePatternParamPartType() == RulePatternParamPartTypeFixed
}

func IsRulePatternParamPartTypeVariadic(p RulePatternParamPart) bool {
	return p.RulePatternParamPartType() == RulePatternParamPartTypeVariadic
}
