package base

type Pattern interface {
	PatternType() PatternType
}

type PatternType string

const (
	PatternTypeVariable PatternType = "VARIABLE"
	PatternTypeBoolean  PatternType = "BOOLEAN"
	PatternTypeString   PatternType = "STRING"
	PatternTypeNumber   PatternType = "NUMBER"
	PatternTypeClass    PatternType = "CLASS"
	PatternTypeTag      PatternType = "TAG"
	PatternTypeObject   PatternType = "OBJECT"
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

func ToPattern[T Pattern](x T) Pattern { return x }
