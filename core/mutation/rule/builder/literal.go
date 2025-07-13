package builder

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type LiteralBuilderFactory struct {
	coreLiteral nodeBuilderFactory
}

func NewLiteralBuilderFactoryWithClassCollection() LiteralBuilderFactory {
	coreLiteral := newCoreLiteralBuilderFactory()
	return NewLiteralBuilderFactory(coreLiteral)
}

func NewLiteralBuilderFactory(coreLiteral nodeBuilderFactory) LiteralBuilderFactory {
	return LiteralBuilderFactory{coreLiteral: coreLiteral}
}

func (f LiteralBuilderFactory) NewBuilder(r stResult.Node) mutator.Builder {
	return wrapChainRemainingChildren(f.coreLiteral.NewBuilder(r))
}

type coreLiteralBuilderFactory struct {
	nonObject coreNonObjectBuilderFactory
	object    coreObjectBuilderFactory
}

func newCoreLiteralBuilderFactory() *coreLiteralBuilderFactory {
	f := &coreLiteralBuilderFactory{}
	f.nonObject = newCoreNonObjectBuilderFactory(f)
	f.object = newCoreObjectBuilderFactory(f)
	return f
}

func (f *coreLiteralBuilderFactory) NewBuilder(r stResult.Node) mutator.Builder {
	if nonObj, ok := f.nonObject.NewBuilder(r).Return(); ok {
		return nonObj
	} else if stResult.IsNodeTypeObject(r) {
		if obj, ok := f.object.NewBuilder(stResult.UnsafeNodeToObject(r)).Return(); ok {
			return obj
		}
	}
	panic("not implemented")
}
