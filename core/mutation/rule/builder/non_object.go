package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	stBase "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type coreNonObjectBuilderFactory struct {
	constant      constantBuilderFactory
	structure     structureBuilderFactory
	reconstructor reconstructorBuilderFactory
}

func newCoreNonObjectBuilderFactory(nodeFactory nodeBuilderFactory) coreNonObjectBuilderFactory {
	return coreNonObjectBuilderFactory{
		constant:      newConstantBuilderFactory(),
		structure:     newStructureBuilderFactory(nodeFactory),
		reconstructor: newReconstructorBuilderFactory(nodeFactory),
	}
}

func (f coreNonObjectBuilderFactory) NewBuilder(r stResult.Node) optional.Of[mutator.Builder] {
	if constant := f.constant.NewBuilder(r); constant.IsNotEmpty() {
		return constant
	}
	switch {
	case stResult.IsNodeTypeStructure(r):
		return optional.Value[mutator.Builder](f.structure.NewBuilder(stResult.UnsafeNodeToStructure(r)))
	case stResult.IsNodeTypeReconstructor(r):
		return f.reconstructor.NewBuilder(stResult.UnsafeNodeToReconstructor(r))
	case stResult.IsNodeTypeVariable(r):
		return optional.Value[mutator.Builder](newParamVariableBuilder(stBase.UnsafeRuleResultToVariable(r)))
	}
	return optional.Empty[mutator.Builder]()
}
