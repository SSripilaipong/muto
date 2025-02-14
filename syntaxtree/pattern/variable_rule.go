package pattern

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type VariableRule struct {
	variableName string
	params       ParamPart
}

func (VariableRule) RulePatternParamType() base.PatternParamType {
	return base.PatternParamTypeNestedVariableRule
}

func (p VariableRule) VariableName() string {
	return p.variableName
}

func (p VariableRule) Variable() base.Variable {
	return base.NewVariable(p.VariableName())
}

func (p VariableRule) ParamPart() ParamPart {
	return p.params
}

func (p VariableRule) ParamParts() []ParamPart {
	return slc.Pure(p.params)
}

func NewVariableRulePattern(variableName string, params ParamPart) VariableRule {
	return VariableRule{variableName: variableName, params: params}
}

func UnsafeRuleParamPatternToVariableRulePattern(p base.PatternParam) VariableRule {
	return p.(VariableRule)
}

func VariableRulePatternToRulePatternParam(x VariableRule) base.PatternParam {
	return x
}
