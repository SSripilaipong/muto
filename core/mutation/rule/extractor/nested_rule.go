package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newNestedNamedRuleExtractor(p stPattern.NamedRule) func(base.Node) optional.Of[*data.Mutation] {
	extractFromChildren := newWithStrictlyChildrenMatch(p.ParamPart())

	return func(x base.Node) optional.Of[*data.Mutation] {
		switch {
		case base.IsObjectNode(x):
			return extractNestedNamedRuleExtractorForNamedObject(p.ObjectName(), extractFromChildren, base.UnsafeNodeToObject(x))
		case base.IsClassNode(x):
			return extractNestedNamedRuleExtractorForClass(p.ObjectName(), extractFromChildren, base.UnsafeNodeToClass(x))
		}
		return optional.Empty[*data.Mutation]()
	}
}

func extractNestedNamedRuleExtractorForClass(name string, extractFromChildren func(obj base.Object) optional.Of[*data.Mutation], class base.Class) optional.Of[*data.Mutation] {
	if class.Name() != name {
		return optional.Empty[*data.Mutation]()
	}
	return extractFromChildren(base.NewObject(class, nil))
}

func extractNestedNamedRuleExtractorForNamedObject(name string, extractFromChildren func(obj base.Object) optional.Of[*data.Mutation], obj base.Object) optional.Of[*data.Mutation] {
	if !base.IsClassNode(obj.Head()) {
		return optional.Empty[*data.Mutation]()
	}
	if base.UnsafeNodeToClass(obj.Head()).Name() != name {
		return optional.Empty[*data.Mutation]()
	}
	return extractFromChildren(obj)
}

func newNestedVariableRuleExtractor(p stPattern.VariableRule) func(base.Node) optional.Of[*data.Mutation] {
	extractFromChildren := newWithStrictlyChildrenMatch(p.ParamPart())
	name := p.VariableName()

	return func(x base.Node) optional.Of[*data.Mutation] {
		switch {
		case base.IsObjectNode(x):
			obj := base.UnsafeNodeToObject(x)
			if mutation, ok := extractFromChildren(obj).Return(); ok {
				return mutation.WithVariableMappings(data.NewVariableMapping(name, obj.Head()))
			}
		}
		return optional.Empty[*data.Mutation]()
	}
}

func newNestedAnonymousRuleExtractor(p stPattern.AnonymousRule) func(base.Node) optional.Of[*data.Mutation] {
	extractFromChildren := newWithStrictlyChildrenMatch(p.ParamPart())
	extractHead := newParamExtractor(p.Head())
	isWrappedObject := stPattern.IsParamTypeNestedNamedRule(p.Head())

	return func(x base.Node) optional.Of[*data.Mutation] {
		if !base.IsObjectNode(x) {
			return optional.Empty[*data.Mutation]()
		}

		obj := base.UnsafeNodeToObject(x)
		mutation, ok := extractFromChildren(obj).Return()
		if !ok {
			return optional.Empty[*data.Mutation]()
		}

		if isWrappedObject && !base.IsObjectNode(obj.Head()) {
			return optional.Empty[*data.Mutation]()
		}
		return optional.JoinFmap(mutation.Merge)(extractHead(obj.Head()))
	}
}

func newWithStrictlyChildrenMatch(paramPart stPattern.ParamPart) func(obj base.Object) optional.Of[*data.Mutation] {
	return newForParamPart(paramPart, strictlyMatchChildren)
}
