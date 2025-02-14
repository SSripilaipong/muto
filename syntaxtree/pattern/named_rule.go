package pattern

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type NamedRule struct { // TODO replace this with AnonymousRule and rename it
	objectName string
	params     ParamPart
}

func (NamedRule) RulePatternParamType() base.PatternParamType {
	return base.PatternParamTypeNestedNamedRule
}

func (r NamedRule) ObjectName() string {
	return r.objectName
}

func (r NamedRule) Object() base.Class {
	return base.NewClass(r.ObjectName())
}

func (r NamedRule) ParamPart() ParamPart {
	return r.params
}

func (r NamedRule) ParamParts() []ParamPart {
	return slc.Pure(r.params)
}

func NewNamedRule(objectName string, params ParamPart) NamedRule {
	return NamedRule{objectName: objectName, params: params}
}

func NamedRuleToParam(x NamedRule) base.PatternParam {
	return x
}

func UnsafeParamToNamedRule(p base.PatternParam) NamedRule {
	return p.(NamedRule)
}
