package mutator

import (
	"phi-lang/common/fn"
	"phi-lang/common/optional"
	"phi-lang/core/base"
)

func New[T, M any](builder Builder[M], extractor Extractor[T, M]) func(t T) optional.Of[base.Node] {
	return fn.Compose(optional.JoinFmap(builder.Build), extractor.Extract)
}
