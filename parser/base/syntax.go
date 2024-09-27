package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

func WithTrailingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.Map(
		tuple.Fn2(func(r R, _ tk.Token) R { return r }),
		ps.Sequence2(p, ps.ConsumeIf(tk.IsLineBreak)),
	)
}

func WithLeadingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.Map(
		tuple.Fn2(func(_ tk.Token, r R) R { return r }),
		ps.Sequence2(ps.ConsumeIf(tk.IsLineBreak), p),
	)
}

func IgnoreTrailingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.DrainTrailing(tk.IsLineBreak, p)
}

func IgnoreLeadingLineBreak[R any](p func([]tk.Token) []tuple.Of2[R, []tk.Token]) func([]tk.Token) []tuple.Of2[R, []tk.Token] {
	return ps.DrainLeading(tk.IsLineBreak, p)
}

func InParentheses[T any](x func([]tk.Token) []tuple.Of2[T, []tk.Token]) func([]tk.Token) []tuple.Of2[T, []tk.Token] {
	return ps.Map(WithoutParenthesis[T], ps.Sequence3(OpenParenthesis, x, CloseParenthesis))
}

func WithoutParenthesis[T any](x tuple.Of3[tk.Token, T, tk.Token]) T {
	return x.X2()
}
