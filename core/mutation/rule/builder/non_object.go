package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type coreNonObjectBuilderFactory struct {
	constant  constantBuilderFactory
	structure structureBuilderFactory
}

func newCoreNonObjectBuilderFactory(nodeFactory nodeBuilderFactory, classCollection ClassCollection) coreNonObjectBuilderFactory {
	return coreNonObjectBuilderFactory{
		constant:  newConstantBuilderFactory(classCollection),
		structure: newStructureBuilderFactory(nodeFactory),
	}
}

func (f coreNonObjectBuilderFactory) NewBuilder(r stResult.Node) optional.Of[mutator.Builder] {
	if constant := f.constant.NewBuilder(r); constant.IsNotEmpty() {
		return constant
	}
	switch {
	case stResult.IsNodeTypeStructure(r):
		return optional.Value[mutator.Builder](f.structure.NewBuilder(stResult.UnsafeNodeToStructure(r)))
	case stResult.IsNodeTypeVariable(r):
		return optional.Value[mutator.Builder](newParamVariableBuilder(stBase.UnsafeRuleResultToVariable(r)))
	}
	return optional.Empty[mutator.Builder]()
}
