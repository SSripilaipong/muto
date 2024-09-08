package extractor

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func extractChildrenNodes(params []st.RuleParamPattern) func(children []base.Node) optional.Of[*data.Mutation] {
	paramExtract := newParamExtractors(params)
	nConsumed := len(paramExtract)

	return func(children []base.Node) optional.Of[*data.Mutation] {
		if len(params) > len(children) {
			return optional.Empty[*data.Mutation]()
		}

		mutation := data.NewMutation()
		for i, child := range children[:nConsumed] {
			e := paramExtract[i](child)
			if e.IsEmpty() {
				return optional.Empty[*data.Mutation]()
			}
			m := mutation.Merge(e.Value())
			if m.IsEmpty() {
				return optional.Empty[*data.Mutation]()
			}
			mutation = m.Value()
		}
		return optional.Value(mutation.WithRemainingChildren(children[nConsumed:]))
	}
}
