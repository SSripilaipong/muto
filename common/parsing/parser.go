package parsing

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

type Parser[R, S any] func([]S) ParseResult[R, S]

func (p Parser[R, S]) F(xs []S) ParseResult[R, S] {
	return p(xs)
}

func (p Parser[R, S]) Legacy(xs []S) tuple.Of2[rslt.Of[R], []S] {
	r, remaining := p(xs).Return()
	return tuple.New2(r, remaining)
}

func ToParser[R, S any](fn func([]S) tuple.Of2[rslt.Of[R], []S]) Parser[R, S] {
	return func(xs []S) ParseResult[R, S] {
		return ParseResult[R, S](fn(xs))
	}
}

type ParseResult[R, S any] tuple.Of2[rslt.Of[R], []S]

func NewParseResult[R, S any](x rslt.Of[R], remaining []S) ParseResult[R, S] {
	return ParseResult[R, S](tuple.New2(x, remaining))
}

func NewParseResultValue[R, S any](x R, remaining []S) ParseResult[R, S] {
	return NewParseResult(rslt.Value(x), remaining)
}

func NewParseResultError[R, S any](err error, remaining []S) ParseResult[R, S] {
	return NewParseResult(rslt.Error[R](err), remaining)
}

func (p ParseResult[R, S]) IsError() bool {
	return tuple.Of2[rslt.Of[R], []S](p).X1().IsErr()
}

func (p ParseResult[R, S]) IsOk() bool {
	return tuple.Of2[rslt.Of[R], []S](p).X1().IsOk()
}

func (p ParseResult[R, S]) Value() R {
	t := tuple.Of2[rslt.Of[R], []S](p)
	if t.X1().IsErr() {
		var zero R
		return zero
	}
	return t.X1().Value()
}

func (p ParseResult[R, S]) Error() error {
	t := tuple.Of2[rslt.Of[R], []S](p)
	if t.X1().IsOk() {
		return nil
	}
	return t.X1().Error()
}

func (p ParseResult[R, S]) Remaining() []S {
	return tuple.Of2[rslt.Of[R], []S](p).X2()
}

func (p ParseResult[R, S]) Return() (rslt.Of[R], []S) {
	t := tuple.Of2[rslt.Of[R], []S](p)
	return t.X1(), t.X2()
}

func (p ParseResult[R, S]) ReturnAsValue() (R, []S) {
	return p.Value(), p.Remaining()
}

func (p ParseResult[R, S]) ReturnAsError() (error, []S) {
	return p.Error(), p.Remaining()
}
