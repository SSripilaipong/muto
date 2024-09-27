package pattern

type VariadicParamPart interface {
	RulePatternVariadicParamPartType() VariadicParamPartType
}

type VariadicParamPartType string

const (
	VariadicParamPartTypeLeft  VariadicParamPartType = "LEFT"
	VariadicParamPartTypeRight VariadicParamPartType = "RIGHT"
)

func IsVariadicParamPartTypeLeft(pp VariadicParamPart) bool {
	return pp.RulePatternVariadicParamPartType() == VariadicParamPartTypeLeft
}

func IsVariadicParamPartTypeRight(pp VariadicParamPart) bool {
	return pp.RulePatternVariadicParamPartType() == VariadicParamPartTypeRight
}

func UnsafeParamPartToVariadicParamPart(p ParamPart) VariadicParamPart {
	return p.(VariadicParamPart)
}

type LeftVariadicParamPart struct {
	name      string
	otherPart FixedParamPart
}

func (LeftVariadicParamPart) RulePatternParamPartType() ParamPartType {
	return ParamPartTypeVariadic
}

func (LeftVariadicParamPart) RulePatternVariadicParamPartType() VariadicParamPartType {
	return VariadicParamPartTypeLeft
}

func (p LeftVariadicParamPart) OtherPart() FixedParamPart {
	return p.otherPart
}

func (p LeftVariadicParamPart) Name() string {
	return p.name
}

func UnsafeVariadicParamPartToLeftVariadicParamPart(p VariadicParamPart) LeftVariadicParamPart {
	return p.(LeftVariadicParamPart)
}

func NewLeftVariadicParamPart(name string, otherPart FixedParamPart) LeftVariadicParamPart {
	return LeftVariadicParamPart{name: name, otherPart: otherPart}
}

type RightVariadicParamPart struct {
	name      string
	otherPart FixedParamPart
}

func (RightVariadicParamPart) RulePatternParamPartType() ParamPartType {
	return ParamPartTypeVariadic
}

func (RightVariadicParamPart) RulePatternVariadicParamPartType() VariadicParamPartType {
	return VariadicParamPartTypeRight
}

func (p RightVariadicParamPart) OtherPart() FixedParamPart {
	return p.otherPart
}

func (p RightVariadicParamPart) Name() string {
	return p.name
}

func UnsafeVariadicParamPartToRightVariadicParamPart(p VariadicParamPart) RightVariadicParamPart {
	return p.(RightVariadicParamPart)
}

func NewRightVariadicParamPart(name string, otherPart FixedParamPart) RightVariadicParamPart {
	return RightVariadicParamPart{name: name, otherPart: otherPart}
}

type MiddleVariadicParamPart struct {
	name      string
	rightPart ParamPart
	leftPart  ParamPart
}

func (MiddleVariadicParamPart) RulePatternParamPartType() ParamPartType {
	return ParamPartTypeVariadic
}

func NewMiddleVariadicParamPart(name string, leftPart ParamPart) MiddleVariadicParamPart {
	return MiddleVariadicParamPart{name: name, leftPart: leftPart}
}
