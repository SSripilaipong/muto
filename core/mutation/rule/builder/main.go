package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func New(r st.RuleResult) func(*data.Mutation) optional.Of[base.Node] {
	switch {
	case st.IsRuleResultTypeBoolean(r):
		return buildBoolean(st.UnsafeRuleResultToBoolean(r))
	case st.IsRuleResultTypeString(r):
		return buildString(st.UnsafeRuleResultToString(r))
	case st.IsRuleResultTypeNumber(r):
		return buildNumber(st.UnsafeRuleResultToNumber(r))
	case st.IsRuleResultTypeNamedObject(r):
		return buildNamedObject(st.UnsafeRuleResultToNamedObject(r))
	case st.IsRuleResultTypeAnonymousObject(r):
		return buildAnonymousObject(st.UnsafeRuleResultToAnonymousObject(r))
	case st.IsRuleResultTypeVariable(r):
		return buildVariable(st.UnsafeRuleResultToVariable(r))
	}
	panic("not implemented")
}
