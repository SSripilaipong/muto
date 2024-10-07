package base

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
)

type Parser[T any] func([]Character) []tuple.Of2[T, []Character]

func (r Parser[T]) FunctionForm() func([]Character) []tuple.Of2[T, []Character] {
	return r
}

type ParserResult[T any] []tuple.Of2[T, []Character]

func AsParserResult[T any](x []tuple.Of2[T, []Character]) ParserResult[T] {
	return x
}

func EmptyResult[T any]() ParserResult[T] {
	return nil
}

func SingleResult[T any](x T, remaining []Character) ParserResult[T] {
	return slc.Pure(tuple.New2(x, remaining))
}
