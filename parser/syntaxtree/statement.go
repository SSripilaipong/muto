package syntaxtree

type Statement interface{}

type Rule struct {
	pattern RulePattern
	result  RuleResult
}

func NewRule(p RulePattern, r RuleResult) Rule {
	return Rule{pattern: p, result: r}
}

type RulePattern struct {
	objectName string
}

func NewRulePattern(objectName string) RulePattern {
	return RulePattern{objectName: objectName}
}

type RuleResult struct {
	objectName string
}

func NewRuleResult(objectName string) RuleResult {
	return RuleResult{objectName: objectName}
}

func RuleToStatement(r Rule) Statement {
	return r
}
