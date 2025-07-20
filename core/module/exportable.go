package module

import (
	stBase "github.com/SSripilaipong/muto/syntaxtree"
)

type Serializable struct {
	Base
	syntaxTree stBase.Module
}

func NewSerializable(module Base, syntaxTree stBase.Module) Serializable {
	return Serializable{Base: module, syntaxTree: syntaxTree}
}

func (m Serializable) Serializable() Serialized {
	return Serialized{
		mainModule: m.syntaxTree,
	}
}

type Serialized struct {
	mainModule stBase.Module
}

func (m Serialized) MainModule() stBase.Module {
	return m.mainModule
}
