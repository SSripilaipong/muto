package automaton

type LineBreak struct {
	namer
}

func newLineBreak() LineBreak {
	return LineBreak{
		namer: newConstNamer(NameLineBreak),
	}
}

func (t LineBreak) Accept(x rune) (Abstract, bool) {
	if x == '\n' {
		return newLineBreak(), true
	}
	return t, false
}
