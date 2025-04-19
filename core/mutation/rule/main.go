package mutation

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/builder"
	ruleExtractor "github.com/SSripilaipong/muto/core/mutation/rule/extractor"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func New(rule st.Rule) mutator.NameWrapper {
	coreBuilder := builder.New(rule.Result())
	nodeBuilder := fixFreeObject(rule.Pattern(), rule.Result(), coreBuilder)
	return mutator.NewNameWrapper(
		rule.PatternName(),
		mutator.NewReconstructor(ruleExtractor.New(rule.Pattern()), nodeBuilder),
	)
}

func fixFreeObject(pattern stPattern.DeterminantObject, result stResult.SimplifiedNode, b mutator.Builder) mutator.Builder {
	if hasExtraParentheses(pattern) && isNonPrimitiveNakedObject(result) {
		return newFreeObjectBuilderGuard(b)
	}
	return b
}

func isNonPrimitiveNakedObject(result stResult.SimplifiedNode) bool {
	return stResult.IsSimplifiedNodeTypeNakedObject(result) &&
		hasResultParam(stResult.UnsafeSimplifiedNodeToNakedObject(result))
}

func hasExtraParentheses(pattern stPattern.DeterminantObject) bool {
	paramPart := pattern.ParamPart()

	return stPattern.IsParamPartTypeFixed(paramPart) &&
		stPattern.UnsafeParamPartToFixedParamPart(paramPart).Size() == 0
}

func hasResultParam(obj stResult.NakedObject) bool {
	paramPart := obj.ParamPart()
	return stResult.IsParamPartTypeFixed(paramPart) &&
		stResult.UnsafeParamPartToFixedParamPart(paramPart).Size() > 0
}
