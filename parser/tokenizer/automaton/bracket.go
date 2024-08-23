package automaton

type Bracket struct {
	namer
}

func newBracket() Bracket {
	return Bracket{
		namer: newConstNamer(NameBracket),
	}
}

func (b Bracket) Accept(rune) (Abstract, bool) {
	return b, false
}
