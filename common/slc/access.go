package slc

import (
	"github.com/SSripilaipong/muto/common/typ"
)

func LastDefaultZero[T any](xs []T) T {
	if len(xs) == 0 {
		return typ.Zero[T]()
	}
	return xs[LastIndex(xs)]
}
