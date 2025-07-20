package base

type ImportedClass struct {
	Class
	module string
}

func NewImportedClass(module string, class Class) ImportedClass {
	return ImportedClass{
		Class:  class,
		module: module,
	}
}

func NewUnlinkedImportedClass(module, name string) ImportedClass {
	return NewImportedClass(module, NewUnlinkedRuleBasedClass(name))
}

func (c ImportedClass) ClassType() ClassType { return ClassTypeImported }

func (c ImportedClass) Module() string {
	return c.module
}

func UnsafeClassToImportedClass(class Class) ImportedClass {
	return class.(ImportedClass)
}
