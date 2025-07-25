package builder

import (
	"fmt"
	"slices"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type ObjectBuilderFactory struct {
	object coreObjectBuilderFactory
}

func NewObjectBuilderFactory() ObjectBuilderFactory {
	return ObjectBuilderFactory{
		object: newCoreObjectBuilderFactory(newCoreLiteralBuilderFactory()),
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

func newCoreObjectBuilderFactory(nodeFactory nodeBuilderFactory) coreObjectBuilderFactory {
	return coreObjectBuilderFactory{
		head:       newCoreNonObjectBuilderFactory(nodeFactory),
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
	head       mutator.Builder
	paramChain paramChainBuilder
}

func newObjectBuilder(head mutator.Builder, paramChain paramChainBuilder) objectBuilder {
	return objectBuilder{
		head:       head,
		paramChain: paramChain,
	}
}

func (b objectBuilder) Build(param *parameter.Parameter) optional.Of[base.Node] {
	head, hasHead := b.head.Build(param).Return()
	if !hasHead {
		return optional.Empty[base.Node]()
	}

	paramChain, ok := b.paramChain.Build(param).Return()
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

func (b objectBuilder) VisitClass(f func(base.Class)) {
	mutator.VisitClass(f, b.head)
	mutator.VisitClass(f, b.paramChain)
}

func (b objectBuilder) DisplayString() string {
	return fmt.Sprintf("(%s)", b.NakedDisplayString())
}

func (b objectBuilder) NakedDisplayString() string {
	headString := DisplayString(b.head)
	return b.paramChain.WrapDisplayString(headString)
}
