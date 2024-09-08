package syntaxtree

type RulePatternVariadicParamPart interface {
	RulePatternVariadicParamPartType() RulePatternVariadicParamPartType
}

type RulePatternVariadicParamPartType string

const (
	RulePatternVariadicParamPartTypeLeft  RulePatternVariadicParamPartType = "LEFT"
	RulePatternVariadicParamPartTypeRight RulePatternVariadicParamPartType = "RIGHT"
)

func IsRulePatternVariadicParamPartTypeLeft(pp RulePatternVariadicParamPart) bool {
	return pp.RulePatternVariadicParamPartType() == RulePatternVariadicParamPartTypeLeft
}

func IsRulePatternVariadicParamPartTypeRight(pp RulePatternVariadicParamPart) bool {
	return pp.RulePatternVariadicParamPartType() == RulePatternVariadicParamPartTypeRight
}

func UnsafeRulePatternParamPartToVariadicParamPart(p RulePatternParamPart) RulePatternVariadicParamPart {
	return p.(RulePatternVariadicParamPart)
}

type RulePatternLeftVariadicParamPart struct {
	name      string
	otherPart RulePatternFixedParamPart
}

func (RulePatternLeftVariadicParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeVariadic
}

func (RulePatternLeftVariadicParamPart) RulePatternVariadicParamPartType() RulePatternVariadicParamPartType {
	return RulePatternVariadicParamPartTypeLeft
}

func (p RulePatternLeftVariadicParamPart) OtherPart() RulePatternFixedParamPart {
	return p.otherPart
}

func (p RulePatternLeftVariadicParamPart) Name() string {
	return p.name
}

func UnsafeRulePatternVariadicParamPartTypeToLeftVariadic(p RulePatternVariadicParamPart) RulePatternLeftVariadicParamPart {
	return p.(RulePatternLeftVariadicParamPart)
}

func NewRulePatternLeftVariadicParamPart(name string, otherPart RulePatternFixedParamPart) RulePatternLeftVariadicParamPart {
	return RulePatternLeftVariadicParamPart{name: name, otherPart: otherPart}
}

type RulePatternRightVariadicParamPart struct {
	name      string
	otherPart RulePatternFixedParamPart
}

func (RulePatternRightVariadicParamPart) RulePatternParamPartType() RulePatternParamPartType {
	return RulePatternParamPartTypeVariadic
}

func (RulePatternRightVariadicParamPart) RulePatternVariadicParamPartType() RulePatternVariadicParamPartType {
	return RulePatternVariadicParamPartTypeRight
}

func (p RulePatternRightVariadicParamPart) OtherPart() RulePatternFixedParamPart {
	return p.otherPart
}

func (p RulePatternRightVariadicParamPart) Name() string {
	return p.name
}

func UnsafeRulePatternVariadicParamPartTypeToRightVariadic(p RulePatternVariadicParamPart) RulePatternRightVariadicParamPart {
	return p.(RulePatternRightVariadicParamPart)
}

func NewRulePatternRightVariadicParamPart(name string, otherPart RulePatternFixedParamPart) RulePatternRightVariadicParamPart {
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

func NewRulePatternMiddleVariadicParamPart(name string, leftPart RulePatternParamPart) RulePatternMiddleVariadicParamPart {
	return RulePatternMiddleVariadicParamPart{name: name, leftPart: leftPart}
}
