package builder

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type SimplifiedNodeBuilderFactory struct {
	object  coreObjectBuilderFactory
	literal nodeBuilderFactory
}

func NewSimplifiedNodeBuilderFactory() SimplifiedNodeBuilderFactory {
	coreLiteral := newCoreLiteralBuilderFactory()
	return SimplifiedNodeBuilderFactory{
		object:  newCoreObjectBuilderFactory(coreLiteral),
		literal: coreLiteral,
	}
}

func (f SimplifiedNodeBuilderFactory) NewBuilder(r stResult.SimplifiedNode) mutator.Builder { // TODO unit test
	if stResult.IsSimplifiedNodeTypeObject(r) {
		if builder, ok := f.object.NewBuilder(stResult.UnsafeSimplifiedNodeToObject(r)).Return(); ok {
			return wrapAppendRemainingChildren(builder)
		}
	} else if stResult.IsSimplifiedNodeTypeNakedObject(r) {
		obj := stResult.UnsafeSimplifiedNodeToNakedObject(r)
		fixedParams := obj.ParamPart()
		if fixedParams.Size() == 0 {
			return wrapAppendRemainingChildren(f.literal.NewBuilder(obj.Head()))
		}
		if builder, ok := f.object.NewBuilder(obj.AsObject()).Return(); ok {
			return wrapChainRemainingChildren(builder)
		}
	}
	panic("not implemented")
}
