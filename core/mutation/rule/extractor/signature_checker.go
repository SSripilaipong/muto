package extractor

import (
	st "phi-lang/syntaxtree"
)

func newSignatureChecker(pattern st.RulePattern) func(ObjectLike) bool {
	return func(obj ObjectLike) bool {
		children := obj.Children()
		if len(children) < int(pattern.NParams()) {
			return false
		}
		return true
	}
}
