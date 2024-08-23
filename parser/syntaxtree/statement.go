package syntaxtree

type Statement interface{}

type Rule struct {
	pattern RulePattern
	result  RuleResult
}

func NewRule(p RulePattern, r RuleResult) Rule {
	return Rule{pattern: p, result: r}
}

func RuleToStatement(r Rule) Statement {
	return r
}
