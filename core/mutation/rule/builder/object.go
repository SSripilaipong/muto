package builder

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func buildObject(obj stResult.Object) func(*data.Mutation) optional.Of[base.Node] {
	buildHead := New(obj.Head())
	buildChildren := buildChildren(obj.ParamPart())

	return func(mapping *data.Mutation) optional.Of[base.Node] {
		children, ok := buildChildren(mapping).Return()
		if !ok {
			return optional.Empty[base.Node]()
		}
		head, hasHead := buildHead(mapping).Return()
		if !hasHead {
			return optional.Empty[base.Node]()
		}

		return optional.Value[base.Node](base.NewObject(head, children))
	}
}

func buildObjectParam(p stResult.Param) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	switch {
	case stResult.IsParamTypeSingle(p):
		return buildObjectParamTypeSingle(p)
	case stResult.IsParamTypeVariadic(p):
		return buildObjectParamTypeVariadic(p)
	}
	panic("not implemented")
}

func buildObjectParamTypeSingle(p stResult.Param) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	return fn.Compose(
		optional.Map(slc.Pure[base.Node]), New(stResult.ParamToNode(p)),
	)
}

func buildObjectParamTypeVariadic(p stResult.Param) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	name := stResult.UnsafeParamToVariadicVariable(p).Name()

	return func(mapping *data.Mutation) optional.Of[[]base.Node] {
		return mapping.VariadicVarValue(name)
	}
}
