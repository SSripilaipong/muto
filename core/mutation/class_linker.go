package mutation

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

type ClassLinker struct{}

func NewClassLinker() ClassLinker {
	return ClassLinker{}
}

func (lk ClassLinker) GetClassFactory() *ClassCollectionAdapter {
	return newClassCollectionAdapter()
}

func (lk ClassLinker) Link(ruleCollection mutator.RuleCollection) {
	// TODO implement
}
