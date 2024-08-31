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
	buildParams := slc.Map(buildObjectParam)(obj.Params())

	return func(mapping *data.Mutation) optional.Of[base.Node] {
		newObj := base.NewNamedObject(obj.ObjectName(), nil)

		return fn.Compose3(
			optional.Map(base.ObjectToNode),
			slc.Fold(aggregateObjectParams)(optional.Value[base.Object](newObj)),
			slc.Map(fn.CallWith[optional.Of[base.Node]](mapping)),
		)(buildParams)
	}
}

var aggregateObjectParams = optional.MergeFn(func(obj base.Object, n base.Node) optional.Of[base.Object] {
	return optional.Value(base.NewObject(obj.Class(), append(obj.Children(), n)))
})

func buildObjectParam(p st.ObjectParam) func(mapping *data.Mutation) optional.Of[base.Node] {
	return New(st.ObjectParamToRuleResult(p))
}
