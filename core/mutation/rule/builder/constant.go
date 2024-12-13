package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type constantBuilder[T base.Node] struct {
	value T
}

func newConstantBuilder[T base.Node](value T) constantBuilder[T] {
	return constantBuilder[T]{value: value}
}

func (b constantBuilder[T]) Build(_ *parameter.Parameter) optional.Of[base.Node] {
	return optional.Value[base.Node](b.value)
}
