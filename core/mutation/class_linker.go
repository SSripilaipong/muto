package mutation

import (
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type ClassLinker struct {
	classes map[string]*base.Class
}

func NewClassLinker() ClassLinker {
	return ClassLinker{classes: make(map[string]*base.Class)}
}

func (lk ClassLinker) GetClass(name string) *base.Class {
	if class, exists := lk.classes[name]; exists {
		return class
	}
	class := base.NewUnlinkedClass(name)
	lk.classes[name] = class
	return class
}

func (lk ClassLinker) LinkCollection(collection mutator.RuleCollection) {
	for name, m := range collection.IterMutators() {
		lk.Link(name, m)
	}
}

func (lk ClassLinker) Link(name string, m base.Mutator) {
	lk.GetClass(name).Link(m)
}
