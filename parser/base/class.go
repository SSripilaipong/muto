package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokens"
)

var Class = ps.Or(
	consumeId(psPred.IsClassName),
	consumeSymbol(psPred.IsSymbol),
	ps.Filter(validClassName, identifierStartingWithLowerCase),
	ps.Filter(classSymbol, symbol),
)

func validClassName(x tk.Token) bool {
	s := x.Value()
	return !psPred.IsBooleanValue(s)
}

func classSymbol(x tk.Token) bool {
	s := x.Value()
	return !psPred.IsEqualSign(s)
}
