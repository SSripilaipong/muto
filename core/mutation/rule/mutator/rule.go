package mutator

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Rule struct {
	Extractor
	Builder
}

func NewRule(extractor Extractor, builder Builder) Rule {
	return Rule{
		Extractor: extractor,
		Builder:   builder,
	}
}

func (m Rule) Mutate(x base.Object) optional.Of[base.Node] {
	return optional.JoinFmap(m.Build)(m.Extract(x))
}

func (m Rule) VisitClass(visitor ClassVisitor) {
	VisitClass(visitor.Visit, m.Builder)
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
