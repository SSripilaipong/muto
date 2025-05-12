package builder

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type ObjectBuilderFactory struct {
	object coreObjectBuilderFactory
}

func NewObjectBuilderFactory(classCollection ClassCollection) ObjectBuilderFactory {
	return ObjectBuilderFactory{
		object: newCoreObjectBuilderFactory(newCoreLiteralBuilderFactory(classCollection), classCollection),
	}
}

func (f ObjectBuilderFactory) NewBuilder(obj stResult.Object) mutator.Builder {
	return wrapChainRemainingChildren(func() mutator.Builder {
		if builder, ok := f.object.NewBuilder(obj).Return(); ok {
			return builder
		}
		panic("not implemented")
	}())
}

type coreObjectBuilderFactory struct {
	head       coreNonObjectBuilderFactory
	paramChain paramChainBuilderFactory
}

func newCoreObjectBuilderFactory(nodeFactory nodeBuilderFactory, classCollection ClassCollection) coreObjectBuilderFactory {
	return coreObjectBuilderFactory{
		head:       newCoreNonObjectBuilderFactory(nodeFactory, classCollection),
		paramChain: newParamChainBuilderFactory(nodeFactory),
	}
}

func (f coreObjectBuilderFactory) NewBuilder(obj stResult.Object) optional.Of[objectBuilder] {
	var params []stResult.FixedParamPart
	var head stResult.Node = obj
	for stResult.IsNodeTypeObject(head) {
		headObj := stResult.UnsafeNodeToObject(head)
		params = append(params, headObj.ParamPart())
		head = headObj.Head()
	}
	slices.Reverse(params)

	headBuilder, headBuilderOk := f.head.NewBuilder(head).Return()
	if !headBuilderOk {
		return optional.Empty[objectBuilder]()
	}
	return optional.Value(newObjectBuilder(headBuilder, f.paramChain.NewBuilder(params)))
}

type objectBuilder struct {
	buildHead       mutator.Builder
	buildParamChain func(mapping *parameter.Parameter) optional.Of[base.ParamChain]
}

func newObjectBuilder(head mutator.Builder, paramChain func(mapping *parameter.Parameter) optional.Of[base.ParamChain]) objectBuilder {
	return objectBuilder{
		buildHead:       head,
		buildParamChain: paramChain,
	}
}

func (b objectBuilder) Build(param *parameter.Parameter) optional.Of[base.Node] {
	head, hasHead := b.buildHead.Build(param).Return()
	if !hasHead {
		return optional.Empty[base.Node]()
	}

	paramChain, ok := b.buildParamChain(param).Return()
	if !ok {
		return optional.Empty[base.Node]()
	}

	if base.IsObjectNode(head) {
		if base.IsCompoundObject(base.UnsafeNodeToObject(head)) {
			return optional.Value[base.Node](base.UnsafeNodeToObject(head).AppendParams(paramChain))
		}
		panic("not implemented")
	}
	return optional.Value[base.Node](base.NewCompoundObject(head, paramChain))
}
