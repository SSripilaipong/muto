package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/syntaxtree"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type nonObjectFactory struct {
	variable variableParamPartFactory
}

func newNonObjectFactory(variable variableParamPartFactory) nonObjectFactory {
	return nonObjectFactory{variable: variable}
}

func (f nonObjectFactory) TryNonObject(p base.Pattern) optional.Of[extractor.NodeExtractor] {
	switch {
	case base.IsPatternTypeBoolean(p):
		return optional.Value(newBooleanParamExtractor(syntaxtree.UnsafePatternToBoolean(p)))
	case base.IsPatternTypeString(p):
		return optional.Value(newStringParamExtractor(syntaxtree.UnsafePatternToString(p)))
	case base.IsPatternTypeRune(p):
		return optional.Value(newRuneParamExtractor(syntaxtree.UnsafePatternToRune(p)))
	case base.IsPatternTypeNumber(p):
		return optional.Value(newNumberParamExtractor(syntaxtree.UnsafePatternToNumber(p)))
	case base.IsPatternTypeTag(p):
		return optional.Value(newTagParamExtractor(syntaxtree.UnsafePatternToTag(p)))
	case base.IsPatternTypeClass(p):
		return optional.Value(newClassParamExtractor(syntaxtree.UnsafePatternToClass(p)))
	case base.IsPatternTypeVariable(p):
		return optional.Value(f.variable.Variable(syntaxtree.UnsafePatternToVariable(p)))
	}
	return optional.Empty[extractor.NodeExtractor]()
}

func (f nonObjectFactory) NonObject(p base.Pattern) extractor.NodeExtractor {
	if r, ok := f.TryNonObject(p).Return(); ok {
		return r
	}
	panic("not implemented")
}

func newBooleanParamExtractor(v syntaxtree.Boolean) extractor.NodeExtractor {
	return extractor.NewBoolean(v.BooleanValue())
}

func newStringParamExtractor(v syntaxtree.String) extractor.NodeExtractor {
	return extractor.NewString(v.StringValue())
}

func newRuneParamExtractor(v syntaxtree.Rune) extractor.NodeExtractor {
	return extractor.NewRune(v.RuneValue())
}

func newNumberParamExtractor(v syntaxtree.Number) extractor.NodeExtractor {
	return extractor.NewNumber(v.NumberValue())
}

func newTagParamExtractor(v syntaxtree.Tag) extractor.NodeExtractor {
	return extractor.NewTag(v.Name())
}

func newClassParamExtractor(v syntaxtree.Class) extractor.NodeExtractor {
	return extractor.NewClass(v.Name())
}
