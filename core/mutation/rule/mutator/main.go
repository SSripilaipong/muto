package mutator

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

func New[T, M any](builder Builder[M], extractor Extractor[T, M]) func(t T) optional.Of[base.Node] {
	return fn.Compose(optional.JoinFmap(builder.Build), extractor.Extract)
}
