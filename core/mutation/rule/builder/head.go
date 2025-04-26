package builder

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type headBuilderFactory struct {
	constant  constantBuilderFactory
	structure structureBuilderFactory
}

func newHeadBuilderFactory(nodeFactory nodeBuilderFactory, classCollection ClassCollection) headBuilderFactory {
	return headBuilderFactory{
		constant:  newConstantBuilderFactory(classCollection),
		structure: newStructureBuilderFactory(nodeFactory),
	}
}

func (f headBuilderFactory) HeadBuilder(r stResult.Node) mutator.Builder {
	if x, ok := f.constant.NewBuilder(r).Return(); ok {
		return x
	}
	switch {
	case stResult.IsNodeTypeStructure(r):
		return f.structure.NewBuilder(stResult.UnsafeNodeToStructure(r))
	case stResult.IsNodeTypeVariable(r):
		return newHeadVariableBuilder(stBase.UnsafeRuleResultToVariable(r))
	}
	panic("not implemented")
}
