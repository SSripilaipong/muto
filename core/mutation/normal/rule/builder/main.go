package builder

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/normal/rule/data"
	st "muto/syntaxtree"
)

func New(r st.RuleResult) func(*data.Mutation) optional.Of[base.Node] {
	switch {
	case st.IsRuleResultTypeString(r):
		return buildString(r.(st.String))
	case st.IsRuleResultTypeNumber(r):
		return buildNumber(r.(st.Number))
	case st.IsRuleResultTypeNamedObject(r):
		return buildNamedObject(r.(st.RuleResultNamedObject))
	case st.IsRuleResultTypeAnonymousObject(r):
		return buildAnonymousObject(r.(st.RuleResultAnonymousObject))
	case st.IsRuleResultTypeVariable(r):
		return buildVariable(r.(st.Variable))
	}
	panic("not implemented")
}
