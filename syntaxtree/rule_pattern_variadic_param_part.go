package syntaxtree

type RulePatternVariadicParamPart interface {
}

func UnsafeRulePatternParamPartToVariadicParamPart(p RulePatternParamPart) RulePatternVariadicParamPart {
	return p.(RulePatternVariadicParamPart)
}

type RulePatternLeftVariadicParamPart struct {
	name      string
	otherPart RulePatternParamPart
}

func (RulePatternLeftVariadicParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeVariadic
}

func (p RulePatternLeftVariadicParamPart) CheckNParams(n int) bool {
	return p.otherPart.CheckNParams(n)
}

func NewRulePatternLeftVariadicParamPart(name string, otherPart RulePatternParamPart) RulePatternLeftVariadicParamPart {
	return RulePatternLeftVariadicParamPart{name: name, otherPart: otherPart}
}

type RulePatternRightVariadicParamPart struct {
	name      string
	otherPart RulePatternParamPart
}

func (RulePatternRightVariadicParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeVariadic
}

func (p RulePatternRightVariadicParamPart) CheckNParams(n int) bool {
	return p.otherPart.CheckNParams(n)
}

func NewRulePatternRightVariadicParamPart(name string, otherPart RulePatternParamPart) RulePatternRightVariadicParamPart {
	return RulePatternRightVariadicParamPart{name: name, otherPart: otherPart}
}

type RulePatternMiddleVariadicParamPart struct {
	name      string
	rightPart RulePatternParamPart
	leftPart  RulePatternParamPart
}

func (RulePatternMiddleVariadicParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeVariadic
}

func (p RulePatternMiddleVariadicParamPart) CheckNParams(n int) bool {
	return p.rightPart.CheckNParams(n)
}

func NewRulePatternMiddleVariadicParamPart(name string, leftPart RulePatternParamPart) RulePatternMiddleVariadicParamPart {
	return RulePatternMiddleVariadicParamPart{name: name, leftPart: leftPart}
}
