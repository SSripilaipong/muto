package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func New(r stResult.Node) func(*data.Mutation) optional.Of[base.Node] {
	switch {
	case stResult.IsNodeTypeBoolean(r):
		return buildBoolean(st.UnsafeRuleResultToBoolean(r))
	case stResult.IsNodeTypeString(r):
		return buildString(st.UnsafeRuleResultToString(r))
	case stResult.IsNodeTypeNumber(r):
		return buildNumber(st.UnsafeRuleResultToNumber(r))
	case stResult.IsNodeTypeClass(r):
		return buildClass(st.UnsafeRuleResultToClass(r))
	case stResult.IsNodeTypeTag(r):
		return buildTag(st.UnsafeRuleResultToTag(r))
	case stResult.IsNodeTypeObject(r):
		return buildObject(stResult.UnsafeNodeToObject(r))
	case stResult.IsNodeTypeVariable(r):
		return buildVariable(st.UnsafeRuleResultToVariable(r))
	}
	panic("not implemented")
}
