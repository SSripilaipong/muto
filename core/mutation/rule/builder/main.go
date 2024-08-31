package builder

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func New(r st.RuleResult) func(*data.Mutation) optional.Of[base.Node] {
	if st.IsRuleResultTypeString(r) {
		return buildString(r.(st.String))
	} else if st.IsRuleResultTypeNumber(r) {
		return buildNumber(r.(st.Number))
	} else if st.IsRuleResultTypeObject(r) {
		return buildObject(r.(st.RuleResultObject))
	} else if st.IsRuleResultTypeVariable(r) {
		return buildVariable(r.(st.Variable))
	}
	panic("not implemented")
}
