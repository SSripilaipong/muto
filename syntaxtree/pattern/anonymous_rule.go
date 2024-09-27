package pattern

type AnonymousRule struct {
	head   Param
	params ParamPart
}

func (AnonymousRule) RulePatternParamType() ParamType {
	return ParamTypeNestedAnonymousRule
}

func (p AnonymousRule) Head() Param {
	return p.head
}

func (p AnonymousRule) ParamPart() ParamPart {
	return p.params
}

func NewAnonymousRule(head Param, params ParamPart) AnonymousRule {
	return AnonymousRule{head: head, params: params}
}

func UnsafeParamToAnonymousRule(p Param) AnonymousRule {
	return p.(AnonymousRule)
}

func AnonymousRuleToParam(x AnonymousRule) Param {
	return x
}
