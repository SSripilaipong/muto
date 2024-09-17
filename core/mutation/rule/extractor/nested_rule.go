package extractor

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func newNestedNamedRuleExtractor(p st.NamedRulePattern) func(base.Node) optional.Of[*data.Mutation] {
	extractFromChildren := newWithStrictlyChildrenMatch(p.ParamPart())

	return func(x base.Node) optional.Of[*data.Mutation] {
		if base.IsNamedObjectNode(x) && base.UnsafeNodeToNamedObject(x).Name() == p.ObjectName() {
			return extractFromChildren(base.UnsafeNodeToObject(x))
		}
		return optional.Empty[*data.Mutation]()
	}
}

func newNestedVariableRuleExtractor(p st.VariableRulePattern) func(base.Node) optional.Of[*data.Mutation] {
	extractFromChildren := newWithStrictlyChildrenMatch(p.ParamPart())
	name := p.VariableName()

	return func(x base.Node) optional.Of[*data.Mutation] {
		switch {
		case base.IsNamedObjectNode(x):
			obj := base.UnsafeNodeToNamedObject(x)
			if mutation, ok := extractFromChildren(obj).Return(); ok {
				return mutation.WithVariableMappings(data.NewVariableMapping(name, base.NewNamedObject(obj.Name(), nil)))
			}
		case base.IsAnonymousObjectNode(x):
			obj := base.UnsafeNodeToAnonymousObject(x)
			if mutation, ok := extractFromChildren(obj).Return(); ok {
				return mutation.WithVariableMappings(data.NewVariableMapping(name, obj.Head()))
			}
		}
		return optional.Empty[*data.Mutation]()
	}
}

func newNestedAnonymousRuleExtractor(p st.AnonymousRulePattern) func(base.Node) optional.Of[*data.Mutation] {
	extractFromChildren := newWithStrictlyChildrenMatch(p.ParamPart())
	extractHead := newParamExtractor(p.Head())

	return func(x base.Node) optional.Of[*data.Mutation] {
		if !base.IsAnonymousObjectNode(x) {
			return optional.Empty[*data.Mutation]()
		}
		obj := base.UnsafeNodeToAnonymousObject(x)

		mutation, ok := extractFromChildren(obj).Return()
		if !ok {
			return optional.Empty[*data.Mutation]()
		}

		return optional.JoinFmap(mutation.Merge)(extractHead(obj.Head()))
	}
}

func newWithStrictlyChildrenMatch(paramPart st.RulePatternParamPart) func(obj base.Object) optional.Of[*data.Mutation] {
	return newForParamPart(paramPart, strictlyMatchChildren)
}