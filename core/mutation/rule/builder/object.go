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

func newObjectBuilder(obj stResult.Object) objectBuilder {
	var params []stResult.ParamPart
	var head stResult.Node = obj
	for stResult.IsNodeTypeObject(head) {
		headObj := stResult.UnsafeNodeToObject(head)
		params = append(params, headObj.ParamPart())
		head = headObj.Head()
	}
	slices.Reverse(params)

	return objectBuilder{
		buildHead:       buildNonObject(head),
		buildParamChain: buildParamChain(params),
	}
}

func (b objectBuilder) Build(param *parameter.Parameter) optional.Of[base.Node] {
	paramChain, ok := b.buildParamChain(param).Return()
	if !ok {
		return optional.Empty[base.Node]()
	}
	head, hasHead := b.buildHead.Build(param).Return()
	if !hasHead {
		return optional.Empty[base.Node]()
	}

	return optional.Value[base.Node](base.NewObject(head, paramChain))
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
		optional.Map(slc.Pure[base.Node]), New(stResult.UnsafeParamToNode(p)).Build,
	)
}

func buildObjectParamTypeVariadic(p stResult.Param) func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
	name := stResult.UnsafeParamToVariadicVariable(p).Name()

	return func(mapping *parameter.Parameter) optional.Of[[]base.Node] {
		return mapping.VariadicVarValue(name)
	}
}
