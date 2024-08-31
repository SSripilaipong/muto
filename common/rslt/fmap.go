package rslt

import "muto/common/fn"

func Fmap[A, B any](f func(A) B) func(Of[A]) Of[B] {
	return func(x Of[A]) Of[B] {
		if x.IsErr() {
			return Error[B](x.Error())
		}
		return Value(f(x.Value()))
	}
}

func Join[T any](x Of[Of[T]]) Of[T] {
	v, err := x.Return()
	if err != nil {
		return Error[T](err)
	}
	return v
}

func JoinFmap[A, B any](f func(A) Of[B]) func(Of[A]) Of[B] {
	return fn.Compose(Join[B], Fmap(f))
}
