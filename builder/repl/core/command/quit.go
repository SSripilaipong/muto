package command

type Quit struct {
	TypeMixin
}

func NewQuit() Quit {
	return Quit{TypeMixin: NewTypeMixin(TypeQuit)}
}
