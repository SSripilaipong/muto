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
		return base.NewNamedObject(name, children)
	}

	return fn.Compose(optional.Map(buildObject), buildChildren(obj.ParamPart()))
}

func buildAnonymousObject(obj st.RuleResultAnonymousObject) func(*data.Mutation) optional.Of[base.Node] {
	buildHead := New(obj.Head())
	buildParams := slc.Map(buildObjectParam)(obj.Params())

	return func(mapping *data.Mutation) optional.Of[base.Node] {
		head, hasHead := buildHead(mapping).Return()
		if !hasHead {
			return optional.Empty[base.Node]()
		}

		newObj := base.NewAnonymousObject(head, nil)
		return fn.Compose3(
			optional.Map(base.AnonymousObjectToNode),
			slc.Fold(aggregateAnonymousObjectParams)(optional.Value[base.AnonymousObject](newObj)),
			slc.Map(fn.CallWith[optional.Of[base.Node]](mapping)),
		)(buildParams)
	}
}

var aggregateAnonymousObjectParams = optional.MergeFn(func(obj base.AnonymousObject, n base.Node) optional.Of[base.AnonymousObject] {
	return optional.Value(base.NewAnonymousObject(obj.Head(), append(obj.Children(), n)))
})

func buildObjectParam(p st.ObjectParam) func(mapping *data.Mutation) optional.Of[base.Node] {
	return New(st.ObjectParamToRuleResult(p))
}
