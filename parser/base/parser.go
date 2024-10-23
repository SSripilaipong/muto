package base

import (
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
)

type Parser[T any] func([]Character) []tuple.Of2[T, []Character]

func (r Parser[T]) FunctionForm() func([]Character) []tuple.Of2[T, []Character] { return r }

type RsParser[T any] func([]Character) []tuple.Of2[rslt.Of[T], []Character]

func (r RsParser[T]) FunctionForm() func([]Character) []tuple.Of2[rslt.Of[T], []Character] { return r }

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

func IgnoreLineAndColumnInResult[T any](x []tuple.Of2[T, []Character]) []tuple.Of2[T, []Character] {
	return slc.Map(tuple.Of2MapX2[T](IgnoreLineAndColumn))(x)
}

func IgnoreLineAndColumn(tokens []Character) []Character {
	var result []Character
	for i := range tokens {
		result = append(result, tokens[i].
			ReplaceColumnNumber(0).
			ReplaceLineNumber(0))
	}
	return result
}
