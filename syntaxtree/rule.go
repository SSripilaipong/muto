package syntaxtree

type Rule struct {
	pattern RulePattern
	result  RuleResult
}

func NewRule(p RulePattern, r RuleResult) Rule {
	return Rule{pattern: p, result: r}
}

func (r Rule) StatementType() StatementType { return RuleStatement }

func (r Rule) Result() RuleResult { return r.result }

func (r Rule) Pattern() RulePattern { return r.pattern }

func (r Rule) PatternName() string { return r.Pattern().ObjectName() }

func RuleToStatement(r Rule) Statement { return r }

func RuleToPatternName(r Rule) string {
	return r.PatternName()
}

func UnsafeStatementToRule(s Statement) Rule {
	return s.(Rule)
}