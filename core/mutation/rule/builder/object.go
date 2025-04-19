package builder

import (
	"slices"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type objectBuilder struct {
	buildHead       mutator.Builder
	buildParamChain func(mapping *parameter.Parameter) optional.Of[base.ParamChain]
}

func NewObject(obj stResult.Object) mutator.Builder {
	return wrapChainRemainingChildren(func() mutator.Builder {
		if builder, ok := newCoreObject(obj).Return(); ok {
			return builder
		}
		panic("not implemented")
	}())
}

func newCoreObject(obj stResult.Object) optional.Of[objectBuilder] {
	var params []stResult.ParamPart
	var head stResult.Node = obj
	for stResult.IsNodeTypeObject(head) {
		headObj := stResult.UnsafeNodeToObject(head)
		params = append(params, headObj.ParamPart())
		head = headObj.Head()
	}
	slices.Reverse(params)

	return optional.Value(objectBuilder{
		buildHead:       buildHead(head),
		buildParamChain: buildParamChain(params),
	})
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

func buildHead(r stResult.Node) mutator.Builder {
	if x, ok := switchConstant(r).Return(); ok {
		return x
	}
	if stResult.IsNodeTypeVariable(r) {
		return newHeadVariableBuilder(stBase.UnsafeRuleResultToVariable(r))
	}
	panic("not implemented")
}

func buildParamChain(paramParts []stResult.ParamPart) func(mapping *parameter.Parameter) optional.Of[base.ParamChain] {
	var childrenBuilders []func(mapping *parameter.Parameter) optional.Of[[]base.Node]
	for _, paramPart := range paramParts {
		childrenBuilders = append(childrenBuilders, buildChildren(paramPart))
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

func buildChildren(paramPart stResult.ParamPart) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	switch {
	case paramPart == nil:
		return func(*parameter.Parameter) optional.Of[[]base.Node] { return optional.Value[[]base.Node](nil) }
	case stResult.IsParamPartTypeFixed(paramPart):
		return buildFixedParamPart(stResult.UnsafeParamPartToFixedParamPart(paramPart))
	}
	panic("not implemented")
}

func buildFixedParamPart(part stResult.FixedParamPart) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	buildParams := slc.Map(buildObjectParam)(part)

	return func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
		return fn.Compose(
			slc.Fold(aggregateNodes)(optional.Value[[]base.Node](nil)),
			slc.Map(fn.CallWith[optional.Of[[]base.Node]](mapping)),
		)(buildParams)
	}
}

func buildObjectParam(p stResult.Param) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	switch {
	case stResult.IsParamTypeSingle(p):
		return buildObjectParamTypeSingle(p)
	case stResult.IsParamTypeVariadic(p):
		return buildObjectParamTypeVariadic(p)
	}
	panic("not implemented")
}

func buildObjectParamTypeSingle(p stResult.Param) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	return fn.Compose(
		optional.Map(slc.Pure[base.Node]), NewLiteralWithoutCarry(stResult.UnsafeParamToNode(p)).Build,
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
