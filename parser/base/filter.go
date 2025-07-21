package base

import (
	"errors"
	"fmt"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/muto/common/parsing"
)

func FilterStatement[T any](raw []tuple.Of2[rslt.Of[T], []Character]) rslt.Of[T] {
	s := parsing.FilterResult(raw)
	if len(s) == 0 {
		var err error
		if len(raw) == 0 {
			err = errors.New("unknown parsing error")
		} else {
			c := raw[0].X2()[0]
			err = fmt.Errorf("parsing error: unexpected token '%c'", c.Value())
		}
		return rslt.Error[T](err)
	}
	r, k := s[0].Return()
	if len(k) > 0 {
		err := errors.New("unknown parsing error")
		if r.IsErr() {
			err = r.Error()
		}
		return rslt.Error[T](fmt.Errorf("parsing error: %w", err))
	}
	if r.IsErr() {
		return rslt.Error[T](fmt.Errorf("parsing error: %w", r.Error()))
	}
	return r
}
