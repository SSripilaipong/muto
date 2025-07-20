package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func newClassBuilder(x st.Class) Class {
	switch {
	case st.IsClassTypeLocal(x):
		return newClass(base.NewUnlinkedRuleBasedClass(st.UnsafeClassToLocalClass(x).Name()))
	case st.IsClassTypeImported(x):
		c := st.UnsafeClassToImportedClass(x)
		return newClass(base.NewUnlinkedImportedClass(c.Module(), c.Name()))
	}
	panic("not implemented")
}

type Class struct {
	value base.Class
}

func newClass(value base.Class) Class {
	return Class{value: value}
}

func (c Class) Build(_ *parameter.Parameter) optional.Of[base.Node] {
	return optional.Value[base.Node](c.value)
}

func (c Class) VisitClass(f func(class base.Class)) {
	f(c.value)
}

func (c Class) DisplayString() string {
	return c.value.TopLevelString()
}
