package builder

import "github.com/SSripilaipong/muto/core/base"

type StaticClassCollection struct {
}

func NewStaticClassCollection() StaticClassCollection {
	return StaticClassCollection{}
}

func (StaticClassCollection) GetClass(name string) *base.Class {
	return base.NewUnlinkedClass(name)
}

var _ ClassCollection = StaticClassCollection{}
