package syntaxtree

type File struct {
	statements []Statement
}

func NewFile(statements []Statement) File {
	return File{statements: statements}
}

var _ Node = File{}
