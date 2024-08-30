package extractor

import (
	"phi-lang/core/base"
	st "phi-lang/syntaxtree"
)

func newSignatureChecker(pattern st.RulePattern) func(base.ObjectLike) bool {
	return func(obj base.ObjectLike) bool {
		children := obj.Children()
		if len(children) < int(pattern.NParams()) {
			return false
		}
		return true
	}
}
