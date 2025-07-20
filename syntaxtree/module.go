package syntaxtree

import (
	"github.com/SSripilaipong/muto/common/slc"
)

type Module struct {
	files []File
}

func NewModule(files []File) Module {
	return Module{files: files}
}

func (p Module) Files() []File {
	return p.files
}

func (p Module) ImportNames() []string {
	var result []string
	for _, file := range p.Files() {
		imports := FilterImportFromStatement(file.Statements())
		result = append(result, slc.Map(ImportToJoinedPath)(imports)...)
	}
	return result
}
