package builder

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func New(r stResult.Node) mutator.Builder {
	return wrapWithRemainingChildren(func() mutator.Builder {
		if x, ok := switchConstant(r).Return(); ok {
			return wrapWithObject(x)
		}
		if stResult.IsNodeTypeVariable(r) {
			return newParamVariableBuilder(stBase.UnsafeRuleResultToVariable(r))
		}
		if stResult.IsNodeTypeObject(r) {
			return newObjectBuilder(stResult.UnsafeNodeToObject(r))
		}
		panic("not implemented")
	}())
}

func buildHead(r stResult.Node) mutator.Builder {
	return wrapWithRemainingChildren(func() mutator.Builder {
		if x, ok := switchConstant(r).Return(); ok {
			return x
		}
		if stResult.IsNodeTypeVariable(r) {
			return newHeadVariableBuilder(stBase.UnsafeRuleResultToVariable(r))
		}
		panic("not implemented")
	}())
}

func switchConstant(r stResult.Node) optional.Of[mutator.Builder] {
	switch {
	case stResult.IsNodeTypeBoolean(r):
		return optional.Value[mutator.Builder](newBooleanBuilder(stBase.UnsafeRuleResultToBoolean(r)))
	case stResult.IsNodeTypeString(r):
		return optional.Value[mutator.Builder](newStringBuilder(stBase.UnsafeRuleResultToString(r)))
	case stResult.IsNodeTypeNumber(r):
		return optional.Value[mutator.Builder](newNumberBuilder(stBase.UnsafeRuleResultToNumber(r)))
	case stResult.IsNodeTypeClass(r):
		return optional.Value[mutator.Builder](newClassBuilder(stBase.UnsafeRuleResultToClass(r)))
	case stResult.IsNodeTypeTag(r):
		return optional.Value[mutator.Builder](newTagBuilder(stBase.UnsafeRuleResultToTag(r)))
	case stResult.IsNodeTypeStructure(r):
		return optional.Value[mutator.Builder](newStructureBuilder(stResult.UnsafeNodeToStructure(r)))
	}
	return optional.Empty[mutator.Builder]()
}

type objectWrapper struct {
	builder mutator.Builder
}

func wrapWithObject(builder mutator.Builder) objectWrapper {
	return objectWrapper{builder: builder}
}

func (o objectWrapper) Build(parameter *parameter.Parameter) optional.Of[base.Node] {
	return optional.Fmap(fn.Compose(base.ToNode, base.NewPrimitiveObject[base.Node]))(o.builder.Build(parameter))
}

var _ mutator.Builder = objectWrapper{}
