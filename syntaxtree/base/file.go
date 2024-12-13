package base

type File struct {
	statements []Statement
}

func NewFile(statements []Statement) File {
	return File{statements: statements}
}

func (f File) Statements() []Statement {
	return f.statements
}
