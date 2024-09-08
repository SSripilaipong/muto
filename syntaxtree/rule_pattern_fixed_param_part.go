package syntaxtree

type RulePatternFixedParamPart []RuleParamPattern

func (RulePatternFixedParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeFixed
}

func (xs RulePatternFixedParamPart) CheckNParams(n int) bool {
	return len(xs) <= n
}

func UnsafeRulePatternParamPartToArrayOfRuleParamPatterns(p RulePatternParamPart) []RuleParamPattern {
	return p.(RulePatternFixedParamPart)
}
