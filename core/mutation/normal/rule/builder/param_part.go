package builder

import (
	"muto/common/fn"
	"muto/common/optional"
	"muto/common/slc"
	"muto/common/tuple"
	"muto/core/base"
	"muto/core/mutation/normal/rule/data"
	st "muto/syntaxtree"
)

func buildChildren(paramPart st.ObjectParamPart) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	switch {
	case paramPart == nil:
		return func(*data.Mutation) optional.Of[[]base.Node] { return optional.Value[[]base.Node](nil) }
	case st.IsObjectParamPartTypeFixed(paramPart):
		return buildFixedParamPart(st.UnsafeObjectParamPartToObjectFixedParamPart(paramPart))
	case st.IsObjectParamPartTypeLeftVariadic(paramPart):
		return buildLeftVariadicParamPart(st.UnsafeObjectParamPartToObjectLeftVariadicParamPart(paramPart))
	case st.IsObjectParamPartTypeRightVariadic(paramPart):
		return buildRightVariadicParamPart(st.UnsafeObjectParamPartToObjectRightVariadicParamPart(paramPart))
	}
	panic("not implemented")
}

func buildLeftVariadicParamPart(param st.ObjectLeftVariadicParamPart) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	build := buildVariadicParamPart(param.Name(), param.OtherPart())

	return func(mapping *data.Mutation) optional.Of[[]base.Node] {
		params := build(mapping)
		if params.IsEmpty() {
			return optional.Empty[[]base.Node]()
		}
		fixedPart, variadicPart := params.Value().Return()
		return optional.Value(append(variadicPart, fixedPart...))
	}
}

func buildRightVariadicParamPart(param st.ObjectRightVariadicParamPart) func(mapping *data.Mutation) optional.Of[[]base.Node] {
	build := buildVariadicParamPart(param.Name(), param.OtherPart())

	return func(mapping *data.Mutation) optional.Of[[]base.Node] {
		params := build(mapping)
		if params.IsEmpty() {
			return optional.Empty[[]base.Node]()
		}
		fixedPart, variadicPart := params.Value().Return()
		return optional.Value(append(fixedPart, variadicPart...))
	}
}

func buildVariadicParamPart(name string, otherPart st.ObjectFixedParamPart) func(mapping *data.Mutation) optional.Of[tuple.Of2[[]base.Node, []base.Node]] {
	buildFixedPart := buildFixedParamPart(otherPart)

	return func(mapping *data.Mutation) optional.Of[tuple.Of2[[]base.Node, []base.Node]] {
		fixedPart, ok := buildFixedPart(mapping).Return()
		if !ok {
			return optional.Empty[tuple.Of2[[]base.Node, []base.Node]]()
		}
		variadicPart, ok := mapping.VariadicVarValue(name).Return()
		if !ok {
			return optional.Empty[tuple.Of2[[]base.Node, []base.Node]]()
		}
		return optional.Value(tuple.New2(fixedPart, variadicPart))
	}
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
