package builder

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type paramChainBuilderFactory struct {
	node nodeBuilderFactory
}

func newParamChainBuilderFactory(nodeFactory nodeBuilderFactory) paramChainBuilderFactory {
	return paramChainBuilderFactory{node: nodeFactory}
}

func (f paramChainBuilderFactory) NewBuilder(paramParts []stResult.FixedParamPart) paramChainBuilder {
	var childrenBuilders []mutator.ListBuilder
	for _, paramPart := range paramParts {
		childrenBuilders = append(childrenBuilders, f.buildChildren(paramPart))
	}
	return paramChainBuilder{childrenBuilders: childrenBuilders}
}

func (f paramChainBuilderFactory) buildChildren(paramPart stResult.FixedParamPart) mutator.ListBuilder {
	if paramPart.Size() == 0 {
		return emptyParamPartBuilder{}
	}
	return fixedParamPartBuilder{buildParams: slc.Map(f.buildObjectParam)(paramPart)}
}

func (f paramChainBuilderFactory) buildObjectParam(p stResult.Param) mutator.ListBuilder {
	switch {
	case stResult.IsParamTypeSingle(p):
		return f.buildObjectParamTypeSingle(stResult.UnsafeParamToNode(p))
	case stResult.IsParamTypeVariadic(p):
		return buildObjectParamTypeVariadic(stResult.UnsafeParamToVariadicVariable(p))
	}
	panic("not implemented")
}

func (f paramChainBuilderFactory) buildObjectParamTypeSingle(param stResult.Node) singleObjectParamBuilder {
	return singleObjectParamBuilder{builder: f.node.NewBuilder(param)}
}

type paramChainBuilder struct {
	childrenBuilders []mutator.ListBuilder
}

func (b paramChainBuilder) Build(mapping *parameter.Parameter) optional.Of[base.ParamChain] {
	var chain [][]base.Node
	for _, childBuilder := range b.childrenBuilders {
		child, ok := childBuilder.Build(mapping).Return()
		if !ok {
			return optional.Empty[base.ParamChain]()
		}
		chain = append(chain, child)
	}
	return optional.Value(base.NewParamChain(chain))
}

func (b paramChainBuilder) WrapDisplayString(head string) string {
	result := head
	for i, builder := range b.childrenBuilders {
		if i > 0 {
			result = fmt.Sprintf("(%s)", result)
		}
		if s := DisplayString(builder); len(s) > 0 {
			result += " " + s
		}
	}
	return result
}

type fixedParamPartBuilder struct {
	buildParams []mutator.ListBuilder
}

func (b fixedParamPartBuilder) Build(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	return fn.Compose(
		slc.Fold(aggregateNodes)(optional.Value[[]base.Node](nil)),
		slc.Map(fn.CallWith[optional.Of[[]base.Node]](mapping)),
	)(slc.Map(mutator.ListBuilderToFunc)(b.buildParams))
}

func (b fixedParamPartBuilder) DisplayString() string {
	var result []string
	for _, param := range b.buildParams {
		result = append(result, DisplayString(param))
	}
	return strings.Join(result, " ")
}

type emptyParamPartBuilder struct{}

func (b emptyParamPartBuilder) Build(*parameter.Parameter) optional.Of[[]base.Node] {
	return optional.Value[[]base.Node](nil)
}

func (b emptyParamPartBuilder) DisplayString() string {
	return ""
}

type singleObjectParamBuilder struct {
	builder mutator.Builder
}

func (b singleObjectParamBuilder) Build(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	return optional.Fmap(slc.Pure[base.Node])(b.builder.Build(mapping))
}

func (b singleObjectParamBuilder) DisplayString() string {
	return DisplayString(b.builder)
}

type variadicObjectParamBuilder struct {
	name string
}

func (b variadicObjectParamBuilder) Build(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	return mapping.VariadicVarValue(b.name)
}

func (b variadicObjectParamBuilder) DisplayString() string {
	return b.name + "..."
}

func buildObjectParamTypeVariadic(p stResult.VariadicVariable) variadicObjectParamBuilder {
	return variadicObjectParamBuilder{name: p.Name()}
}

var aggregateNodes = optional.MergeFn(func(nodes []base.Node, xs []base.Node) optional.Of[[]base.Node] {
	return optional.Value(append(nodes, xs...))
})
