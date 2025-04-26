package mutation

import (
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/builder"
)

type ClassCollectionAdapter struct {
}

func newClassCollectionAdapter() *ClassCollectionAdapter {
	return &ClassCollectionAdapter{}
}

func (*ClassCollectionAdapter) GetClass(name string) base.Class {
	return base.NewClass(name)
}

var _ builder.ClassCollection = &ClassCollectionAdapter{}
