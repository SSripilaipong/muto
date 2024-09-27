package pattern

type NamedRule struct {
	objectName string
	params     ParamPart
}

func (NamedRule) RulePatternParamType() ParamType {
	return ParamTypeNestedNamedRule
}

func (r NamedRule) ObjectName() string {
	return r.objectName
}

func (r NamedRule) ParamPart() ParamPart {
	return r.params
}

func NewNamedRule(objectName string, params ParamPart) NamedRule {
	return NamedRule{objectName: objectName, params: params}
}

func NamedRuleToParam(x NamedRule) Param {
	return x
}

func UnsafeParamToNamedRule(p Param) NamedRule {
	return p.(NamedRule)
}
