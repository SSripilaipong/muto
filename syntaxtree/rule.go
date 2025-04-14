package syntaxtree

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Rule struct {
	pattern stPattern.DeterminantObject
	result  stResult.Node
}

func NewRule(p stPattern.DeterminantObject, r stResult.Node) Rule {
	return Rule{pattern: p, result: r}
}

func (r Rule) StatementType() base.StatementType { return base.RuleStatement }

func (r Rule) Result() stResult.Node { return r.result }

func (r Rule) Pattern() stPattern.DeterminantObject { return r.pattern }

func (r Rule) PatternName() string { return r.Pattern().ObjectName() }

func RuleToStatement(r Rule) base.Statement { return r }

func RuleToPatternName(r Rule) string {
	return r.PatternName()
}

func UnsafeStatementToRule(s base.Statement) Rule {
	return s.(Rule)
}

type ActiveRule struct {
	Rule
}

func (r ActiveRule) StatementType() base.StatementType { return base.ActiveRuleStatement }

func NewActiveRule(p stPattern.DeterminantObject, r stResult.Node) ActiveRule {
	return ActiveRule{Rule: NewRule(p, r)}
}

func ActiveRuleToStatement(r ActiveRule) base.Statement {
	return r
}

func UnsafeStatementToActiveRule(s base.Statement) ActiveRule {
	return s.(ActiveRule)
}
