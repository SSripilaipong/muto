package base

type Import struct {
	path []string
}

func NewImport(path []string) Import {
	return Import{path: path}
}

func (l Import) StatementType() StatementType { return ImportStatement }

func (l Import) Path() []string {
	return l.path
}

func ImportToStatement(l Import) Statement {
	return l
}
