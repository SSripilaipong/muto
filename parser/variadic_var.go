package parser

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
	return rslt.Error[variadicVarNode](errors.New("not a variadic variable"))
})

type variadicVarNode struct {
	name string
}

func (v variadicVarNode) Name() string {
	return v.name
}

func newVariadicVar(name string) variadicVarNode {
	return variadicVarNode{name: name}
}
