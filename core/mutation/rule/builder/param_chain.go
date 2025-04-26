package builder

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type paramChainBuilderFactory struct {
	node nodeBuilderFactory
}

func newParamChainBuilderFactory(nodeFactory nodeBuilderFactory) paramChainBuilderFactory {
	return paramChainBuilderFactory{node: nodeFactory}
}

func (f paramChainBuilderFactory) NewBuilder(paramParts []stResult.FixedParamPart) func(mapping *parameter.Parameter) optional.Of[base.ParamChain] {
	var childrenBuilders []func(mapping *parameter.Parameter) optional.Of[[]base.Node]
	for _, paramPart := range paramParts {
		childrenBuilders = append(childrenBuilders, f.buildChildren(paramPart))
	}
	return func(mapping *parameter.Parameter) optional.Of[base.ParamChain] {
		var chain [][]base.Node
		for _, childBuilder := range childrenBuilders {
			child, ok := childBuilder(mapping).Return()
			if !ok {
				return optional.Empty[base.ParamChain]()
			}
			chain = append(chain, child)
		}
		return optional.Value(base.NewParamChain(chain))
	}
}

func (f paramChainBuilderFactory) buildChildren(paramPart stResult.FixedParamPart) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	if paramPart.Size() == 0 {
		return func(*parameter.Parameter) optional.Of[[]base.Node] { return optional.Value[[]base.Node](nil) }
	}
	return f.buildFixedParamPart(paramPart)
}

func (f paramChainBuilderFactory) buildFixedParamPart(part stResult.FixedParamPart) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	buildParams := slc.Map(f.buildObjectParam)(part)

	return func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
		return fn.Compose(
			slc.Fold(aggregateNodes)(optional.Value[[]base.Node](nil)),
			slc.Map(fn.CallWith[optional.Of[[]base.Node]](mapping)),
		)(buildParams)
	}
}

func (f paramChainBuilderFactory) buildObjectParam(p stResult.Param) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	switch {
	case stResult.IsParamTypeSingle(p):
		return f.buildObjectParamTypeSingle(p)
	case stResult.IsParamTypeVariadic(p):
		return buildObjectParamTypeVariadic(p)
	}
	panic("not implemented")
}

func (f paramChainBuilderFactory) buildObjectParamTypeSingle(p stResult.Param) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	return fn.Compose(
		optional.Map(slc.Pure[base.Node]), f.node.NewBuilder(stResult.UnsafeParamToNode(p)).Build,
	)
}

func buildObjectParamTypeVariadic(p stResult.Param) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	name := stResult.UnsafeParamToVariadicVariable(p).Name()

	return func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
		return mapping.VariadicVarValue(name)
	}
}

var aggregateNodes = optional.MergeFn(func(nodes []base.Node, xs []base.Node) optional.Of[[]base.Node] {
	return optional.Value(append(nodes, xs...))
})
