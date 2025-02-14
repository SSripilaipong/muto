package pattern

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type AnonymousRule struct {
	head   base.PatternParam
	params ParamPart
}

func (AnonymousRule) RulePatternParamType() base.PatternParamType {
	return base.PatternParamTypeNestedAnonymousRule
}

func (p AnonymousRule) Head() base.PatternParam {
	return p.head
}

func (p AnonymousRule) ParamPart() ParamPart {
	return p.params
}

func (p AnonymousRule) ParamParts() []ParamPart {
	return slc.Pure(p.params)
}

func NewAnonymousRule(head base.PatternParam, params ParamPart) AnonymousRule {
	return AnonymousRule{head: head, params: params}
}

func UnsafeParamToAnonymousRule(p base.PatternParam) AnonymousRule {
	return p.(AnonymousRule)
}

func AnonymousRuleToParam(x AnonymousRule) base.PatternParam {
	return x
}
