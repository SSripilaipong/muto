package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func New(r stResult.Node) mutator.Builder {
	return wrapWithRemainingChildren(func() mutator.Builder {
		if x, ok := switchNonObject(r).Return(); ok {
			return x
		}
		if stResult.IsNodeTypeObject(r) {
			return newObjectBuilder(stResult.UnsafeNodeToObject(r))
		}
		panic("not implemented")
	}())
}

func buildNonObject(r stResult.Node) mutator.Builder {
	x, ok := switchNonObject(r).Return()
	if !ok {
		panic("not implemented")
	}
	return wrapWithRemainingChildren(x)
}

func switchNonObject(r stResult.Node) optional.Of[mutator.Builder] {
	switch {
	case stResult.IsNodeTypeBoolean(r):
		return optional.Value[mutator.Builder](newBooleanBuilder(base.UnsafeRuleResultToBoolean(r)))
	case stResult.IsNodeTypeString(r):
		return optional.Value[mutator.Builder](newStringBuilder(base.UnsafeRuleResultToString(r)))
	case stResult.IsNodeTypeNumber(r):
		return optional.Value[mutator.Builder](newNumberBuilder(base.UnsafeRuleResultToNumber(r)))
	case stResult.IsNodeTypeClass(r):
		return optional.Value[mutator.Builder](newClassBuilder(base.UnsafeRuleResultToClass(r)))
	case stResult.IsNodeTypeTag(r):
		return optional.Value[mutator.Builder](newTagBuilder(base.UnsafeRuleResultToTag(r)))
	case stResult.IsNodeTypeStructure(r):
		return optional.Value[mutator.Builder](newStructureBuilder(stResult.UnsafeNodeToStructure(r)))
	case stResult.IsNodeTypeVariable(r):
		return optional.Value[mutator.Builder](newVariableBuilder(base.UnsafeRuleResultToVariable(r)))
	}
	return optional.Empty[mutator.Builder]()
}
