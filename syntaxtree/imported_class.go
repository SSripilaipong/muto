package syntaxtree

type ImportedClass struct {
	LocalClass
	module string
}

func NewImportedClass(module string, name string) ImportedClass {
	return ImportedClass{LocalClass: NewLocalClass(name), module: module}
}

func (ImportedClass) ClassType() ClassType { return ClassTypeImported }

func (c ImportedClass) Module() string { return c.module }

func UnsafeClassToImportedClass(p Class) ImportedClass { return p.(ImportedClass) }
