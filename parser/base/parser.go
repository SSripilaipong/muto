package base

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

type Parser[T any] func([]tk.Token) []tuple.Of2[T, []tk.Token]

func (r Parser[T]) FunctionForm() func([]tk.Token) []tuple.Of2[T, []tk.Token] {
	return r
}

type ParserResult[T any] []tuple.Of2[T, []tk.Token]

func AsParserResult[T any](x []tuple.Of2[T, []tk.Token]) ParserResult[T] {
	return x
}

func EmptyResult[T any]() ParserResult[T] {
	return nil
}

func SingleResult[T any](x T, remaining []tk.Token) ParserResult[T] {
	return slc.Pure(tuple.New2(x, remaining))
}
