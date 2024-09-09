package syntaxtree

type RulePatternFixedParamPart []RuleParamPattern

func (RulePatternFixedParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeFixed
}

func RuleParamPatternsToRulePatternFixedParamPart(xs []RuleParamPattern) RulePatternFixedParamPart {
	return xs
}

func RuleParamPatternsToRulePatternParamPart(xs []RuleParamPattern) RulePatternParamPart {
	return RuleParamPatternsToRulePatternFixedParamPart(xs)
}

func UnsafeRulePatternParamPartToArrayOfRuleParamPatterns(p RulePatternParamPart) []RuleParamPattern {
	return p.(RulePatternFixedParamPart)
}
