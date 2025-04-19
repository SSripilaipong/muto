package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

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
