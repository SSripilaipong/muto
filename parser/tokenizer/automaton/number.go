package automaton

import "unicode"

type Number struct {
	namer
}

func newNumber() Number {
	return Number{
		namer: newConstNamer(NameNumber),
	}
}

func (n Number) Accept(x rune) (Abstract, bool) {
	if unicode.IsDigit(x) {
		return newNumber(), true
	} else if x == '.' {
		return newNumberAfterDot(), true
	}
	return n, false
}

type NumberAfterDot struct {
	namer
}

func newNumberAfterDot() Number {
	return Number{
		namer: newConstNamer(NameNumber),
	}
}

func (n NumberAfterDot) Accept(x rune) (Abstract, bool) {
	if unicode.IsDigit(x) {
		return newNumberAfterDot(), true
	}
	return nil, false
}
