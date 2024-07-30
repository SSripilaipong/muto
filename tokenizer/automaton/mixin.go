package automaton

type namer interface {
	Name() Name
}

type namerFunc func() Name

func (f namerFunc) Name() Name {
	return f()
}

func newConstNamer(name Name) namer { return namerFunc(func() Name { return name }) }
