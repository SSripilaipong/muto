package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func New(r stResult.SimplifiedNode) mutator.Builder { // TODO unit test
	if stResult.IsSimplifiedNodeTypeObject(r) {
		if builder, ok := newCoreObject(stResult.UnsafeSimplifiedNodeToObject(r)).Return(); ok {
			return wrapAppendRemainingChildren(builder)
		}
	} else if stResult.IsSimplifiedNodeTypeNakedObject(r) {
		obj := stResult.UnsafeSimplifiedNodeToNakedObject(r)
		params := obj.ParamPart()
		if stResult.IsParamPartTypeFixed(params) {
			fixedParams := stResult.UnsafeParamPartToFixedParamPart(params)
			if fixedParams.Size() == 0 {
				return NewLiteral(obj.Head())
			}
			if builder, ok := newCoreObject(obj.AsObject()).Return(); ok {
				return wrapChainRemainingChildren(builder)
			}
		}
	}
	panic("not implemented")
}

func NewLiteral(r stResult.Node) mutator.Builder {
	return wrapChainRemainingChildren(func() mutator.Builder {
		if nonObj, ok := newNonObject(r).Return(); ok {
			return nonObj
		} else if stResult.IsNodeTypeObject(r) {
			if obj, ok := newCoreObject(stResult.UnsafeNodeToObject(r)).Return(); ok {
				return obj
			}
		}
		panic("not implemented")
	}())
}

func NewNonObject(r stResult.Node) mutator.Builder {
	return wrapChainRemainingChildren(func() mutator.Builder {
		if builder := newNonObject(r); builder.IsNotEmpty() {
			return builder.Value()
		}
		panic("not implemented")

	}())
}

func newNonObject(r stResult.Node) optional.Of[mutator.Builder] {
	if constant := switchConstant(r); constant.IsNotEmpty() {
		return constant
	}
	if stResult.IsNodeTypeVariable(r) {
		return optional.Value[mutator.Builder](newParamVariableBuilder(stBase.UnsafeRuleResultToVariable(r)))
	}
	return optional.Empty[mutator.Builder]()
}

func buildHead(r stResult.Node) mutator.Builder {
	return wrapChainRemainingChildren(func() mutator.Builder {
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
