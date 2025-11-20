package base

import (
	"errors"
	"fmt"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

func FilterResult[T any](s tuple.Of2[rslt.Of[T], []Character]) rslt.Of[T] {
	r, k := s.Return()
	if len(k) > 0 {
		err := errors.New("unknown parsing error")
		if r.IsErr() {
			err = r.Error()
		}
		c := k[0]
		return rslt.Error[T](fmt.Errorf("parsing error at line %d, column %d: %w", c.LineNumber(), c.ColumnNumber(), err))
	}
	if r.IsErr() {
		return rslt.Error[T](fmt.Errorf("parsing error: %w", r.Error()))
	}
	return r
}
