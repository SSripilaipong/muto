package mutation

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func fixFreeObject(pattern stPattern.DeterminantObject, result stResult.SimplifiedNode, b mutator.Builder) mutator.Builder {
	if hasExtraParentheses(pattern) && isNonPrimitiveNakedObject(result) {
		return newFreeObjectBuilderGuard(b)
	}
	return b
}

func isNonPrimitiveNakedObject(result stResult.SimplifiedNode) bool {
	return stResult.IsSimplifiedNodeTypeNakedObject(result) &&
		stResult.UnsafeSimplifiedNodeToNakedObject(result).ParamPart().Size() > 0
}

func hasExtraParentheses(pattern stPattern.DeterminantObject) bool {
	paramPart := pattern.ParamPart()

	return stPattern.IsParamPartTypeFixed(paramPart) &&
		stPattern.UnsafeParamPartToFixedParamPart(paramPart).Size() == 0
}

type freeObjectBuilderGuard struct {
	builder mutator.Builder
}

func newFreeObjectBuilderGuard(builder mutator.Builder) mutator.Builder {
	return freeObjectBuilderGuard{builder: builder}
}

func (g freeObjectBuilderGuard) Build(params *parameter.Parameter) optional.Of[base.Node] {
	if params.RemainingParamChain().Size() < 1 {
		return optional.Empty[base.Node]()
	}
	return g.builder.Build(params)
}

func (g freeObjectBuilderGuard) VisitClass(f func(base.Class)) {
	mutator.VisitClass(f, g.builder)
}
