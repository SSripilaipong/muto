package extractor

import (
	"muto/core/base"
	st "muto/syntaxtree"
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
