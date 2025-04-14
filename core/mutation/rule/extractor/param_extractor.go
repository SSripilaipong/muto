package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newParamExtractors(params []base.Pattern) []extractor.NodeExtractor {
	return slc.Map(newPatternExtractor)(params)
}

func newPatternExtractor(p base.Pattern) extractor.NodeExtractor {
	if r, ok := tryNonObjectPatternExtractor(p).Return(); ok {
		return r
	}
	if base.IsPatternTypeObject(p) {
		return newObjectExtractor(stPattern.UnsafePatternToObject(p))
	}
	panic("not implemented")
}

func newNonObjectPatternExtractor(p base.Pattern) extractor.NodeExtractor {
	if r, ok := tryNonObjectPatternExtractor(p).Return(); ok {
		return r
	}
	panic("not implemented")
}

func tryNonObjectPatternExtractor(p base.Pattern) optional.Of[extractor.NodeExtractor] {
	switch {
	case base.IsPatternTypeBoolean(p):
		return optional.Value(newBooleanParamExtractor(base.UnsafePatternToBoolean(p)))
	case base.IsPatternTypeString(p):
		return optional.Value(newStringParamExtractor(base.UnsafePatternToString(p)))
	case base.IsPatternTypeNumber(p):
		return optional.Value(newNumberParamExtractor(base.UnsafePatternToNumber(p)))
	case base.IsPatternTypeTag(p):
		return optional.Value(newTagParamExtractor(base.UnsafePatternToTag(p)))
	case base.IsPatternTypeClass(p):
		return optional.Value(newClassParamExtractor(base.UnsafePatternToClass(p)))
	case base.IsPatternTypeVariable(p):
		return optional.Value(newVariableParamExtractor(base.UnsafePatternToVariable(p)))
	}
	return optional.Empty[extractor.NodeExtractor]()
}

func newVariableParamExtractor(v base.Variable) extractor.NodeExtractor {
	return extractor.NewParamVariable(v.Name())
}

func newBooleanParamExtractor(v base.Boolean) extractor.NodeExtractor {
	return extractor.NewBoolean(v.BooleanValue())
}

func newStringParamExtractor(v base.String) extractor.NodeExtractor {
	return extractor.NewString(v.StringValue())
}

func newNumberParamExtractor(v base.Number) extractor.NodeExtractor {
	return extractor.NewNumber(v.NumberValue())
}

func newTagParamExtractor(v base.Tag) extractor.NodeExtractor {
	return extractor.NewTag(v.Name())
}

func newClassParamExtractor(v base.Class) extractor.NodeExtractor {
	return extractor.NewClass(v.Name())
}
