package base

import (
	"errors"
	"strings"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	"github.com/SSripilaipong/muto/parser/tokenizer"
)

var VariadicVar = ps.Transform(func(x tokenizer.Token) rslt.Of[VariadicVarNode] {
	name := x.Value()
	if tokenizer.IsIdentifier(x) && psPred.IsVariadicVariable(name) {
		return rslt.Value(newVariadicVar(strings.Trim(name, ".")))
	}
	return rslt.Error[VariadicVarNode](errors.New("not a variadic variable"))
})

type VariadicVarNode struct {
	name string
}

func (v VariadicVarNode) Name() string {
	return v.name
}

func newVariadicVar(name string) VariadicVarNode {
	return VariadicVarNode{name: name}
}
