package optional

import "phi-lang/common/fn"

func Fmap[A, B any](f func(A) B) func(Of[A]) Of[B] {
	return func(x Of[A]) Of[B] {
		if x.IsEmpty() {
			return Empty[B]()
		}
		return Value(f(x.Value()))
	}
}

func Join[T any](x Of[Of[T]]) Of[T] {
	v, ok := x.Return()
	if !ok {
		return Empty[T]()
	}
	return v
}

func JoinFmap[A, B any](f func(A) Of[B]) func(Of[A]) Of[B] {
	return fn.Compose(Join[B], Fmap(f))
}
