package automaton

import "unicode"

type Symbol struct {
	namer
}

func newSymbol() Symbol {
	return Symbol{
		namer: newConstNamer(NameSymbol),
	}
}

func (s Symbol) Accept(x rune) (Abstract, bool) {
	if unicode.IsSymbol(x) || unicode.IsPunct(x) {
		return newSymbol(), true
	}
	return s, false
}
