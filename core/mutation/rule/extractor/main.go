package extractor

import (
	"phi-lang/common/optional"
	"phi-lang/core/mutation/rule/data"
	st "phi-lang/syntaxtree"
)

func New(rule st.RulePattern) func(obj ObjectLike) optional.Of[*data.Mutation] {
	signatureCheck := newSignatureChecker(rule)
	paramExtract := newParamExtractors(rule)

	return func(obj ObjectLike) optional.Of[*data.Mutation] {
		if !signatureCheck(obj) {
			return optional.Empty[*data.Mutation]()
		}

		mutation := data.NewMutation()
		for i, child := range obj.Children() {
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
		return optional.Value(mutation)
	}
}
