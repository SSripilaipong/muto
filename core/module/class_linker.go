package module

import (
	"maps"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type ClassLinker interface {
	GetOrCreateClass(name string) *base.Class
	LinkCollection(collection mutator.RuleCollection)
	Link(name string, m base.Mutator)
	Clone() ClassLinker
	Mapping() map[string]*base.Class
	ReplaceMapping(mapping map[string]*base.Class)
}

type ClassLinkerImpl struct {
	classes map[string]*base.Class
}

func NewClassLinker() *ClassLinkerImpl {
	return NewClassLinkerFromMapping(make(map[string]*base.Class))
}

func NewClassLinkerFromMapping(classes map[string]*base.Class) *ClassLinkerImpl {
	return &ClassLinkerImpl{classes: classes}
}

func (lk *ClassLinkerImpl) GetOrCreateClass(name string) *base.Class {
	if class, exists := lk.classes[name]; exists {
		return class
	}
	class := base.NewUnlinkedClass(name)
	lk.classes[name] = class
	return class
}

func (lk *ClassLinkerImpl) LinkCollection(collection mutator.RuleCollection) {
	for name, m := range collection.IterMutators() {
		lk.Link(name, m)
	}
}

func (lk *ClassLinkerImpl) Link(name string, m base.Mutator) {
	lk.GetOrCreateClass(name).Link(m)
}

func (lk *ClassLinkerImpl) AddOverwrite(t ClassLinker) ClassLinker {
	mapping := maps.Clone(lk.Mapping())
	maps.Copy(mapping, t.Mapping())
	return NewClassLinkerFromMapping(mapping)
}

func (lk *ClassLinkerImpl) Clone() ClassLinker {
	return NewClassLinkerFromMapping(lk.Mapping())
}

func (lk *ClassLinkerImpl) Mapping() map[string]*base.Class {
	return maps.Clone(lk.classes)
}

func (lk *ClassLinkerImpl) ReplaceMapping(mapping map[string]*base.Class) {
	maps.Copy(lk.classes, mapping)
}
