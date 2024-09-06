package builder

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func buildObject(obj st.RuleResultObject) func(*data.Mutation) optional.Of[base.Node] {
	return buildNamedObject(obj)
}

func buildNamedObject(obj st.RuleResultObject) func(*data.Mutation) optional.Of[base.Node] {
	buildParams := slc.Map(buildObjectParam)(obj.Params())

	return func(mapping *data.Mutation) optional.Of[base.Node] {
		newObj := base.NewNamedObject(obj.ObjectName(), nil)

		return fn.Compose3(
			optional.Map(base.NamedObjectToNode),
			slc.Fold(aggregateObjectParams)(optional.Value[base.NamedObject](newObj)),
			slc.Map(fn.CallWith[optional.Of[base.Node]](mapping)),
		)(buildParams)
	}
}

var aggregateObjectParams = optional.MergeFn(func(obj base.NamedObject, n base.Node) optional.Of[base.NamedObject] {
	return optional.Value(base.NewNamedObject(obj.ClassName(), append(obj.Children(), n)))
})

func buildObjectParam(p st.ObjectParam) func(mapping *data.Mutation) optional.Of[base.Node] {
	return New(st.ObjectParamToRuleResult(p))
}
