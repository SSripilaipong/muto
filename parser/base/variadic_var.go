package base

import (
	"strings"

	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

var VariadicVar = ps.Map(newVariadicVar, ps.Or(
	ps.Map(tk.TokenToValue, ps.Filter(fn.And(tk.IsIdentifier, fn.Compose(psPred.IsVariadicVariable, tk.TokenToValue)), ps.One[tk.Token])),
	ps.Map(tuple.Fn2(joinTokenString), ps.Sequence2(identifierStartingWithUpperCase, fixedChars("..."))),
))

type VariadicVarNode struct {
	name string
}

func (v VariadicVarNode) Name() string {
	return v.name
}

func newVariadicVar(name string) VariadicVarNode {
	return VariadicVarNode{name: strings.Trim(name, ".")}
}
