package base

import (
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Rule struct {
	pattern stPattern.NamedRule
	result  stResult.Node
}

func NewRule(p stPattern.NamedRule, r stResult.Node) Rule {
	return Rule{pattern: p, result: r}
}

func (r Rule) StatementType() StatementType { return RuleStatement }

func (r Rule) Result() stResult.Node { return r.result }

func (r Rule) Pattern() stPattern.NamedRule { return r.pattern }

func (r Rule) PatternName() string { return r.Pattern().ObjectName() }

func RuleToStatement(r Rule) Statement { return r }

func RuleToPatternName(r Rule) string {
	return r.PatternName()
}

func UnsafeStatementToRule(s Statement) Rule {
	return s.(Rule)
}

type ActiveRule struct {
	Rule
}

func (r ActiveRule) StatementType() StatementType { return ActiveRuleStatement }

func NewActiveRule(p stPattern.NamedRule, r stResult.Node) ActiveRule {
	return ActiveRule{Rule{p, r}}
}

func ActiveRuleToStatement(r ActiveRule) Statement {
	return r
}

func UnsafeStatementToActiveRule(s Statement) ActiveRule {
	return s.(ActiveRule)
}
