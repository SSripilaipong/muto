package base

type Pattern interface {
	PatternType() PatternType
}

type PatternType string

const (
	PatternTypeVariable    PatternType = "VARIABLE"
	PatternTypeBoolean     PatternType = "BOOLEAN"
	PatternTypeString      PatternType = "STRING"
	PatternTypeRune        PatternType = "RUNE"
	PatternTypeNumber      PatternType = "NUMBER"
	PatternTypeClass       PatternType = "CLASS"
	PatternTypeTag         PatternType = "TAG"
	PatternTypeObject      PatternType = "OBJECT"
	PatternTypeConjunction PatternType = "CONJUNCTION"
)

func IsPatternTypeVariable(p Pattern) bool {
	return p.PatternType() == PatternTypeVariable
}

func IsPatternTypeBoolean(p Pattern) bool {
	return p.PatternType() == PatternTypeBoolean
}

func IsPatternTypeString(p Pattern) bool {
	return p.PatternType() == PatternTypeString
}

func IsPatternTypeRune(p Pattern) bool {
	return p.PatternType() == PatternTypeRune
}

func IsPatternTypeNumber(p Pattern) bool {
	return p.PatternType() == PatternTypeNumber
}

func IsPatternTypeTag(p Pattern) bool {
	return p.PatternType() == PatternTypeTag
}

func IsPatternTypeClass(p Pattern) bool {
	return p.PatternType() == PatternTypeClass
}

func IsPatternTypeObject(p Pattern) bool {
	return p.PatternType() == PatternTypeObject
}

func IsPatternTypeConjunction(p Pattern) bool {
	return p.PatternType() == PatternTypeConjunction
}

func ToPattern[T Pattern](x T) Pattern { return x }
