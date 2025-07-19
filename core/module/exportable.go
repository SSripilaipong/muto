package module

import stBase "github.com/SSripilaipong/muto/syntaxtree/base"

type Exportable struct {
	Base
	syntaxTree stBase.Module
}

func NewExportable(module Base, syntaxTree stBase.Module) Exportable {
	return Exportable{Base: module, syntaxTree: syntaxTree}
}

func (m Exportable) Export() Exported {
	return Exported{
		mainModule: m.syntaxTree,
	}
}

type Exported struct {
	mainModule stBase.Module
}

func (m Exported) MainModule() stBase.Module {
	return m.mainModule
}
