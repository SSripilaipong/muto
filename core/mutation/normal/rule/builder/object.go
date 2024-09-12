package builder

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/normal/rule/data"
	st "muto/syntaxtree"
)

func buildNamedObject(obj st.RuleResultNamedObject) func(*data.Mutation) optional.Of[base.Node] {
	name := obj.ObjectName()
	buildObject := func(children []base.Node) base.Node {
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

		return optional.Value[base.Node](base.NewAnonymousObject(head, children))
	}
}

var aggregateAnonymousObjectParams = optional.MergeFn(func(obj base.AnonymousObject, n base.Node) optional.Of[base.AnonymousObject] {
	return optional.Value(base.NewAnonymousObject(obj.Head(), append(obj.Children(), n)))
})

func buildObjectParam(p st.ObjectParam) func(mapping *data.Mutation) optional.Of[base.Node] {
	return New(st.ObjectParamToRuleResult(p))
}
