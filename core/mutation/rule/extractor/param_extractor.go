package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/syntaxtree"
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
		return optional.Value(newBooleanParamExtractor(syntaxtree.UnsafePatternToBoolean(p)))
	case base.IsPatternTypeString(p):
		return optional.Value(newStringParamExtractor(syntaxtree.UnsafePatternToString(p)))
	case base.IsPatternTypeNumber(p):
		return optional.Value(newNumberParamExtractor(syntaxtree.UnsafePatternToNumber(p)))
	case base.IsPatternTypeTag(p):
		return optional.Value(newTagParamExtractor(syntaxtree.UnsafePatternToTag(p)))
	case base.IsPatternTypeClass(p):
		return optional.Value(newClassParamExtractor(syntaxtree.UnsafePatternToClass(p)))
	case base.IsPatternTypeVariable(p):
		return optional.Value(newVariableParamExtractor(syntaxtree.UnsafePatternToVariable(p)))
	}
	return optional.Empty[extractor.NodeExtractor]()
}

func newVariableParamExtractor(v syntaxtree.Variable) extractor.NodeExtractor {
	if len(v.Name()) == 0 {
		panic("variable name should not be empty")
	}
	if v.Name()[0] == '_' {
		return extractor.NewIgnoredParamVariable()
	}
	return extractor.NewParamVariable(v.Name())
}

func newBooleanParamExtractor(v syntaxtree.Boolean) extractor.NodeExtractor {
	return extractor.NewBoolean(v.BooleanValue())
}

func newStringParamExtractor(v syntaxtree.String) extractor.NodeExtractor {
	return extractor.NewString(v.StringValue())
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
