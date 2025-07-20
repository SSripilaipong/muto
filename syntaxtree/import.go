package syntaxtree

import (
	"strings"

	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type Import struct {
	path []string
}

func NewImport(path []string) Import {
	return Import{path: path}
}

func (l Import) StatementType() base.StatementType { return base.ImportStatement }

func (l Import) Path() []string {
	return l.path
}

func (l Import) JoinedPath() string {
	return strings.Join(l.path, ".")
}

func ImportToJoinedPath(l Import) string {
	return l.JoinedPath()
}

func ImportToStatement(l Import) base.Statement {
	return l
}

func UnsafeStatementToImport(s base.Statement) Import {
	return s.(Import)
}

var FilterImportFromStatement = fn.Compose(
	slc.Map(UnsafeStatementToImport), slc.Filter(base.IsImportStatement),
)
