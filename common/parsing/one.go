package parsing

import (
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/slc"
)

func One[S any](xs []S) []tuple.Of2[S, []S] {
	if len(xs) == 0 {
		return nil
	}
	return slc.Pure(tuple.New2(xs[0], xs[1:]))
}
