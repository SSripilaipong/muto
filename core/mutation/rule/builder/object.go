package builder

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func buildNamedObject(obj st.RuleResultNamedObject) func(*data.Mutation) optional.Of[base.Node] {
	name := obj.ObjectName()
	buildObject := func(children []base.Node) base.Node {
		if len(children) == 0 {
			return base.NewNamedClass(name)
		}
		return base.NewNamedObject(name, children)
	}

	return fn.Compose(optional.Map(buildObject), buildChildren(obj.ParamPart()))
}

func buildAnonymousObject(obj st.RuleResultAnonymousObject) func(*data.Mutation) optional.Of[base.Node] {
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
		return optional.Value[base.Node](base.NewAnonymousObject(head, children))
	}
}

func autoBubbleUp(head base.Node, children []base.Node) optional.Of[base.Node] {
	if base.IsNamedClassNode(head) {
		return optional.Value[base.Node](base.NewObject(base.UnsafeNodeToNamedClass(head), children))
	}
	return optional.Empty[base.Node]()
}

func buildObjectParam(p st.ObjectParam) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	switch {
	case st.IsObjectParamTypeSingle(p):
		return buildObjectParamTypeSingle(p)
	case st.IsObjectParamTypeVariadic(p):
		return buildObjectParamTypeVariadic(p)
	}
	panic("not implemented")
}

func buildObjectParamTypeSingle(p st.ObjectParam) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	return fn.Compose(
		optional.Map(slc.Pure[base.Node]), New(st.ObjectParamToRuleResult(p)),
	)
}

func buildObjectParamTypeVariadic(p st.ObjectParam) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	name := st.UnsafeObjectParamToVariadicVariable(p).Name()

	return func(mapping *data.Mutation) optional.Of[[]base.Node] {
		return mapping.VariadicVarValue(name)
	}
}
