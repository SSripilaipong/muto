package syntaxtree

type RulePatternFixedParamPart []RuleParamPattern

func (RulePatternFixedParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeFixed
}

func UnsafeRulePatternParamPartToArrayOfRuleParamPatterns(p RulePatternParamPart) []RuleParamPattern {
	return p.(RulePatternFixedParamPart)
}
