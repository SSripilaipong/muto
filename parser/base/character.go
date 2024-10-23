package base

import (
	"fmt"
	"strconv"
)

type Character struct {
	value        rune
	lineNumber   uint
	columnNumber uint
}

func NewCharacter(value rune, lineNumber uint, columnNumber uint) Character {
	return Character{
		value:        value,
		lineNumber:   lineNumber,
		columnNumber: columnNumber,
	}
}

func (t Character) LineNumber() uint {
	return t.lineNumber
}

func (t Character) ColumnNumber() uint {
	return t.columnNumber
}

func (t Character) Value() rune {
	return t.value
}

func (t Character) String() string {
	return fmt.Sprintf("Char(%s)", strconv.QuoteRune(t.value))
}

func (t Character) ReplaceLineNumber(i uint) Character {
	t.lineNumber = i
	return t
}

func (t Character) ReplaceColumnNumber(i uint) Character {
	t.columnNumber = i
	return t
}

func CharacterToValue(t Character) rune {
	return t.Value()
}
