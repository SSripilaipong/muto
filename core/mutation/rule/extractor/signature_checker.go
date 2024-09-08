package extractor

import (
	"muto/core/base"
	st "muto/syntaxtree"
)

func newSignatureChecker(pattern st.RulePattern) func(base.Object) bool {
	return func(obj base.Object) bool {
		return pattern.CheckNParams(len(obj.Children()))
	}
}
