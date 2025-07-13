package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Reconstructor struct {
	Extractor
	Builder
}

func NewReconstructor(extractor Extractor, builder Builder) Reconstructor {
	return Reconstructor{
		Extractor: extractor,
		Builder:   builder,
	}
}

func (m Reconstructor) Mutate(x base.Object) optional.Of[base.Node] {
	return optional.JoinFmap(m.Build)(m.Extract(x))
}

func (m Reconstructor) LinkClass(linker ClassLinker) {
	VisitClass(linker.LinkClass, m.Builder)
}

type Builder interface {
	Build(*parameter.Parameter) optional.Of[base.Node]
}

type ListBuilder interface {
	Build(*parameter.Parameter) optional.Of[[]base.Node]
}

func ListBuilderToFunc(builder ListBuilder) func(*parameter.Parameter) optional.Of[[]base.Node] {
	return builder.Build
}

type Extractor interface {
	Extract(base.Object) optional.Of[*parameter.Parameter]
}
