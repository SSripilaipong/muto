package builder

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func New(r stResult.Node) mutator.Builder {
	return wrapWithRemainingChildren(func() mutator.Builder {
		switch {
		case stResult.IsNodeTypeBoolean(r):
			return newBooleanBuilder(base.UnsafeRuleResultToBoolean(r))
		case stResult.IsNodeTypeString(r):
			return newStringBuilder(base.UnsafeRuleResultToString(r))
		case stResult.IsNodeTypeNumber(r):
			return newNumberBuilder(base.UnsafeRuleResultToNumber(r))
		case stResult.IsNodeTypeClass(r):
			return newClassBuilder(base.UnsafeRuleResultToClass(r))
		case stResult.IsNodeTypeTag(r):
			return newTagBuilder(base.UnsafeRuleResultToTag(r))
		case stResult.IsNodeTypeStructure(r):
			return newStructureBuilder(stResult.UnsafeNodeToStructure(r))
		case stResult.IsNodeTypeObject(r):
			return newObjectBuilder(stResult.UnsafeNodeToObject(r))
		case stResult.IsNodeTypeVariable(r):
			return newVariableBuilder(base.UnsafeRuleResultToVariable(r))
		}
		panic("not implemented")
	}())
}
