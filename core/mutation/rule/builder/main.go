package builder

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	"phi-lang/core/base/datatype"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

func New(r st.RuleResult) func(data.Mutation) optional.Of[base.Node] {
	if st.IsRuleResultTypeString(r) {
		return buildString(r.(st.String))
	} else if st.IsRuleResultTypeNumber(r) {
		return buildNumber(r.(st.Number))
	}
	panic("not implemented")
}

func buildString(s st.String) func(mapping data.Mutation) optional.Of[base.Node] {
	value := s.Value()
	return func(mapping data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](base.NewString(value))
	}
}

func buildNumber(x st.Number) func(data.Mutation) optional.Of[base.Node] {
	value := base.NewNumber(datatype.NewNumber(x.Value()))

	return func(mapping data.Mutation) optional.Of[base.Node] {
		return optional.Value[base.Node](value)
	}
}
