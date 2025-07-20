package syntaxtree

import "github.com/SSripilaipong/muto/syntaxtree/base"

type File struct {
	statements []base.Statement
}

func NewFile(statements []base.Statement) File {
	return File{statements: statements}
}

func (f File) Statements() []base.Statement {
	return f.statements
}
