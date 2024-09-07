package extractor

import (
	"muto/common/optional"
	"muto/core/base"
	"muto/core/mutation/rule/data"
	st "muto/syntaxtree"
)

func New(rule st.RulePattern) func(obj base.Object) optional.Of[*data.Mutation] {
	signatureCheck := newSignatureChecker(rule)
	paramExtract := newParamExtractors(rule)
	nConsumed := len(paramExtract)

	return func(obj base.Object) optional.Of[*data.Mutation] {
		if !signatureCheck(obj) {
			return optional.Empty[*data.Mutation]()
		}

		mutation := data.NewMutation()
		for i, child := range obj.Children()[:nConsumed] {
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
		return optional.Value(mutation.WithRemainingChildren(obj.Children()[nConsumed:]))
	}
}
