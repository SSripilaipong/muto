package extractor

import (
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

type VariableFactory interface {
	FixedVariable(name string) extractor.NodeExtractor
	VariadicVariable(name string) extractor.NodeListExtractor
}

type variableParamPartFactory struct {
	variable VariableFactory
	core     corePatternFactory
}

func newVariableParamPartFactory(core corePatternFactory, variable VariableFactory) variableParamPartFactory {
	return variableParamPartFactory{
		variable: variable,
		core:     core,
	}
}

func (f variableParamPartFactory) VariadicParamPart(paramPart stPattern.VariadicParamPart) extractor.NodeListExtractor {
	switch {
	case stPattern.IsVariadicParamPartTypeRight(paramPart):
		return f.rightVariadic(stPattern.UnsafeVariadicParamPartToRightVariadicParamPart(paramPart))
	case stPattern.IsVariadicParamPartTypeLeft(paramPart):
		return f.leftVariadic(stPattern.UnsafeVariadicParamPartToLeftVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func (f variableParamPartFactory) Variable(v syntaxtree.Variable) extractor.NodeExtractor {
	if len(v.Name()) == 0 {
		panic("variable name should not be empty")
	}
	if v.Name()[0] == '_' {
		return extractor.NewIgnoredParamVariable(v.Name())
	}
	return f.variable.FixedVariable(v.Name())
}

func (f variableParamPartFactory) rightVariadic(pp stPattern.RightVariadicParamPart) extractor.NodeListExtractor {
	if len(pp.Name()) == 0 {
		panic("right variadic param part's name should not be empty")
	}
	return extractor.NewRightVariadic(pp.Name(), f.variadic(pp.Name()), f.core.Patterns(pp.OtherPart()))
}

func (f variableParamPartFactory) leftVariadic(pp stPattern.LeftVariadicParamPart) extractor.NodeListExtractor {
	if len(pp.Name()) == 0 {
		panic("left variadic param part's name should not be empty")
	}
	return extractor.NewLeftVariadic(pp.Name(), f.variadic(pp.Name()), f.core.Patterns(pp.OtherPart()))
}

func (f variableParamPartFactory) variadic(name string) extractor.NodeListExtractor {
	if name[0] == '_' {
		return extractor.NewContextFreeIgnoreVariadic(name)
	}
	return f.variable.VariadicVariable(name)
}
