package automaton

type Name string

const (
	NameEmpty      Name = "EMPTY"
	NameBracket    Name = "BRACKET"
	NameIdentifier Name = "IDENTIFIER"
	NameLineBreak  Name = "LINE_BREAK"
	NameNumber     Name = "NUMBER"
	NameString     Name = "STRING"
	NameSpace      Name = "SPACE"
	NameSymbol     Name = "SYMBOL"
)
