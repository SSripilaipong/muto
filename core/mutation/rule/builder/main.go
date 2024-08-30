package builder

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
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
