package pattern

import (
	"slices"

	"github.com/SSripilaipong/muto/syntaxtree/base"
)

func ExtractNonObjectHead(p base.Pattern) base.Pattern {
	for base.IsPatternTypeObject(p) {
		p = UnsafePatternToObject(p).Head()
	}
	return p
}

func ExtractParamChain(p base.Pattern) []ParamPart { // TODO unit test
	var paramChain []ParamPart // inserted in reversed order first, corrected before return

	for base.IsPatternTypeObject(p) {
		obj := UnsafePatternToObject(p)
		paramChain = append(paramChain, obj.ParamPart())
		p = obj.Head()
	}

	slices.Reverse(paramChain)
	return paramChain
}
