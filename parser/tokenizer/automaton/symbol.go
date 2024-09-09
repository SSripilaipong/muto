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
	if (unicode.IsSymbol(x) || unicode.IsPunct(x)) && !isBracket(x) {
		return newSymbol(), true
	}
	return s, false
}

type LeadingNegative struct {
	namer
}

func newLeadingNegative() LeadingNegative {
	return LeadingNegative{
		namer: newConstNamer(NameSymbol),
	}
}

func (s LeadingNegative) Accept(x rune) (Abstract, bool) {
	if unicode.IsDigit(x) {
		return newNumber(), true
	}
	if (unicode.IsSymbol(x) || unicode.IsPunct(x)) && !isBracket(x) {
		return newSymbol(), true
	}
	return s, false
}
