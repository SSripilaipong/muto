package pattern

import "github.com/SSripilaipong/muto/syntaxtree/base"

type Conjunction struct {
	main base.Pattern
	conj base.Pattern
}

func (Conjunction) PatternType() base.PatternType { return base.PatternTypeConjunction }

func (c Conjunction) Main() base.Pattern {
	return c.main
}

func (c Conjunction) Conj() base.Pattern {
	return c.conj
}

func NewConjunction(main, conj base.Pattern) Conjunction {
	return Conjunction{main: main, conj: conj}
}

func UnsafePatternToConjunction(p base.Pattern) Conjunction {
	return p.(Conjunction)
}
