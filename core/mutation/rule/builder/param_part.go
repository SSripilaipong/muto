package builder

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func buildChildren(paramPart st.ObjectParamPart) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	switch {
	case paramPart == nil:
		return func(*data.Mutation) optional.Of[[]base.Node] { return optional.Value[[]base.Node](nil) }
	case st.IsObjectParamPartTypeFixed(paramPart):
		return buildFixedParamPart(st.UnsafeObjectParamPartToObjectFixedParamPart(paramPart))
	}
	panic("not implemented")
}

func buildFixedParamPart(part st.ObjectFixedParamPart) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	buildParams := slc.Map(buildObjectParam)(part)

	return func(mapping *data.Mutation) optional.Of[[]base.Node] {
		return fn.Compose(
			slc.Fold(aggregateNodes)(optional.Value[[]base.Node](nil)),
			slc.Map(fn.CallWith[optional.Of[base.Node]](mapping)),
		)(buildParams)
	}
}

var aggregateNodes = optional.MergeFn(func(nodes []base.Node, n base.Node) optional.Of[[]base.Node] {
	return optional.Value(append(nodes, n))
})
