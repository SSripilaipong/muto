package syntaxtree

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type ActiveRule struct {
	Rule
}

func (r ActiveRule) StatementType() base.StatementType { return base.ActiveRuleStatement }

func NewActiveRule(p stPattern.DeterminantObject, r stResult.SimplifiedNode) ActiveRule {
	return ActiveRule{Rule: NewRule(p, r)}
}

func ActiveRuleToStatement(r ActiveRule) base.Statement {
	return r
}

func UnsafeStatementToActiveRule(s base.Statement) ActiveRule {
	return s.(ActiveRule)
}

var FilterActiveRuleFromStatement = fn.Compose3(
	slc.Map(activeRuleToRule), slc.Map(UnsafeStatementToActiveRule), slc.Filter(base.IsActiveRuleStatement),
)

func activeRuleToRule(x ActiveRule) Rule {
	return NewRule(x.Pattern(), x.Result())
}
