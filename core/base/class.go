package base

type Class interface {
	Name() string
}

type NamedClass struct {
	name string
}

func NewNamedClass(name string) NamedClass {
	return NamedClass{name: name}
}

func (c NamedClass) Name() string {
	return c.name
}
