package syntaxtree

type Rule struct {
	pattern NamedRulePattern
	result  RuleResult
}

func NewRule(p NamedRulePattern, r RuleResult) Rule {
	return Rule{pattern: p, result: r}
}

func (r Rule) StatementType() StatementType { return RuleStatement }

func (r Rule) Result() RuleResult { return r.result }

func (r Rule) Pattern() NamedRulePattern { return r.pattern }

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

func NewActiveRule(p NamedRulePattern, r RuleResult) ActiveRule {
	return ActiveRule{Rule{p, r}}
}

func ActiveRuleToStatement(r ActiveRule) Statement {
	return r
}
