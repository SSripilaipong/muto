package pattern

import (
	"fmt"
	"slices"

	"github.com/SSripilaipong/muto/syntaxtree/base"
)

func ExtractNonNestedHeadParam(p base.PatternParam) base.PatternParam { // TODO unit test
	for base.IsNestedPatternParam(p) {
		switch {
		case base.IsPatternParamTypeNestedAnonymousRule(p):
			p = UnsafeParamToAnonymousRule(p).Head()
		case base.IsPatternParamTypeNestedNamedRule(p):
			p = UnsafeParamToNamedRule(p).Object()
		case base.IsPatternParamTypeNestedVariableRule(p):
			p = UnsafeRuleParamPatternToVariableRulePattern(p).Variable()
		default:
			panic(fmt.Errorf("unknown nested param type: %#v", p))
		}
	}
	return p
}

func ExtractParamChain(p base.PatternParam) []ParamPart { // TODO unit test
	var paramChain []ParamPart // inserted in reversed order first, corrected before return

	for base.IsNestedPatternParam(p) {
		switch {
		case base.IsPatternParamTypeNestedAnonymousRule(p):
			paramChain = append(paramChain, UnsafeParamToAnonymousRule(p).ParamPart())
			p = UnsafeParamToAnonymousRule(p).Head()
		case base.IsPatternParamTypeNestedNamedRule(p):
			paramChain = append(paramChain, UnsafeParamToNamedRule(p).ParamPart())
			p = UnsafeParamToNamedRule(p).Object()
		case base.IsPatternParamTypeNestedVariableRule(p):
			paramChain = append(paramChain, UnsafeRuleParamPatternToVariableRulePattern(p).ParamPart())
			p = UnsafeRuleParamPatternToVariableRulePattern(p).Variable()
		default:
			panic(fmt.Errorf("unknown nested param type: %#v", p))
		}
	}

	slices.Reverse(paramChain)
	return paramChain
}
