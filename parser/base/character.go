package base

import (
	"fmt"
)

type Character struct {
	value rune
}

func NewCharacter(value rune) Character {
	return Character{value: value}
}

func (t Character) Value() rune {
	return t.value
}

func (t Character) String() string {
	return fmt.Sprintf("character(%#v)", t.value)
}

func CharacterToValue(t Character) rune {
	return t.Value()
}
