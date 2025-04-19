package builder

import (
	"slices"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
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
	paramWithoutRemainingParamChain := param.SetRemainingParamChain(base.NewParamChain(nil))

	head, hasHead := b.buildHead.Build(paramWithoutRemainingParamChain).Return()
	if !hasHead {
		return optional.Empty[base.Node]()
	}

	paramChain, ok := b.buildParamChain(paramWithoutRemainingParamChain).Return()
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
		optional.Map(slc.Pure[base.Node]), NewLiteral(stResult.UnsafeParamToNode(p)).Build,
	)
}

func buildObjectParamTypeVariadic(p stResult.Param) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	name := stResult.UnsafeParamToVariadicVariable(p).Name()

	return func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
		return mapping.VariadicVarValue(name)
	}
}
