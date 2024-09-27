package builder

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func buildNamedObject(obj stResult.NamedObject) func(*data.Mutation) optional.Of[base.Node] {
	name := obj.ObjectName()
	buildObject := func(children []base.Node) base.Node {
		if len(children) == 0 {
			return base.NewClass(name)
		}
		return base.NewNamedObject(name, children)
	}

	return fn.Compose(optional.Map(buildObject), buildChildren(obj.ParamPart()))
}

func buildAnonymousObject(obj stResult.Object) func(*data.Mutation) optional.Of[base.Node] {
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

		if result, ok := autoBubbleUp(head, children).Return(); ok {
			return optional.Value[base.Node](result)
		}
		return optional.Value[base.Node](base.NewObject(head, children))
	}
}

func autoBubbleUp(head base.Node, children []base.Node) optional.Of[base.Node] {
	if base.IsClassNode(head) {
		return optional.Value[base.Node](base.NewObject(base.UnsafeNodeToClass(head), children))
	}
	return optional.Empty[base.Node]()
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
