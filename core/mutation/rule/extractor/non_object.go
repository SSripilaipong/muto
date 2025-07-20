package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	st "github.com/SSripilaipong/muto/syntaxtree"
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
		return optional.Value(newBooleanParamExtractor(st.UnsafePatternToBoolean(p)))
	case base.IsPatternTypeString(p):
		return optional.Value(newStringParamExtractor(st.UnsafePatternToString(p)))
	case base.IsPatternTypeRune(p):
		return optional.Value(newRuneParamExtractor(st.UnsafePatternToRune(p)))
	case base.IsPatternTypeNumber(p):
		return optional.Value(newNumberParamExtractor(st.UnsafePatternToNumber(p)))
	case base.IsPatternTypeTag(p):
		return optional.Value(newTagParamExtractor(st.UnsafePatternToTag(p)))
	case base.IsPatternTypeClass(p):
		return optional.Value(newClassParamExtractor(st.UnsafePatternToClass(p)))
	case base.IsPatternTypeVariable(p):
		return optional.Value(f.variable.Variable(st.UnsafePatternToVariable(p)))
	}
	return optional.Empty[extractor.NodeExtractor]()
}

func (f nonObjectFactory) NonObject(p base.Pattern) extractor.NodeExtractor {
	if r, ok := f.TryNonObject(p).Return(); ok {
		return r
	}
	panic("not implemented")
}

func newBooleanParamExtractor(v st.Boolean) extractor.NodeExtractor {
	return extractor.NewBoolean(v.BooleanValue())
}

func newStringParamExtractor(v st.String) extractor.NodeExtractor {
	return extractor.NewString(v.StringValue())
}

func newRuneParamExtractor(v st.Rune) extractor.NodeExtractor {
	return extractor.NewRune(v.RuneValue())
}

func newNumberParamExtractor(v st.Number) extractor.NodeExtractor {
	return extractor.NewNumber(v.NumberValue())
}

func newTagParamExtractor(v st.Tag) extractor.NodeExtractor {
	return extractor.NewTag(v.Name())
}

func newClassParamExtractor(v st.Class) extractor.NodeExtractor {
	switch {
	case st.IsClassTypeLocal(v):
		local := st.UnsafeClassToLocalClass(v)
		return extractor.NewLocalClass(local.Name())
	case st.IsClassTypeImported(v):
		imported := st.UnsafeClassToImportedClass(v)
		return extractor.NewImportedClass(imported.Module(), imported.Name()) // TODO map module name from alias first (pass mapping in)
	}
	panic("not implemented")
}
