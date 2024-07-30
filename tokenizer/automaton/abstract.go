package automaton

type Abstract interface {
	Accept(x rune) (Abstract, bool)
	Name() Name
}
