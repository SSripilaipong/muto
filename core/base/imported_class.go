package base

import "fmt"

type ImportedClass struct {
	*RuleBasedClass
	module string
}

func NewImportedClass(module string, class *RuleBasedClass) ImportedClass {
	return ImportedClass{
		RuleBasedClass: class,
		module:         module,
	}
}

func NewUnlinkedImportedClass(module, name string) ImportedClass {
	return NewImportedClass(module, NewUnlinkedRuleBasedClass(name))
}

func (c ImportedClass) ClassType() ClassType { return ClassTypeImported }

func (c ImportedClass) Module() string {
	return c.module
}

func (c ImportedClass) TopLevelString() string {
	return c.String()
}

func (c ImportedClass) String() string {
	return fmt.Sprintf("%s.%s", c.Module(), c.Name())
}

func (c ImportedClass) MutoString() string {
	return c.Name()
}

func (c ImportedClass) Equals(d ImportedClass) bool {
	return c.Name() == d.Name() && c.Module() == d.Module() // TODO make sure the module name is absolute, not just alias
}

func UnsafeClassToImportedClass(class Class) ImportedClass {
	return class.(ImportedClass)
}
