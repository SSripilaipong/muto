package mutator

import "phi-lang/common/optional"

type Extractor[T, M any] interface {
	Extract(T) optional.Of[M]
}

func ExtractorFunc[T, M any](f extractorFunc[T, M]) Extractor[T, M] {
	return f
}

type extractorFunc[T, M any] func(T) optional.Of[M]

func (f extractorFunc[T, M]) Extract(x T) optional.Of[M] {
	return f(x)
}
