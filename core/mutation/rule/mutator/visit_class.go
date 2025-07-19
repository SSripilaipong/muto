package mutator

import (
	"github.com/SSripilaipong/muto/core/base"
)

type ClassVisitable interface {
	VisitClass(f func(class base.Class))
}

func VisitClass(f func(class base.Class), x any) {
	if visitable, isVisitable := x.(ClassVisitable); isVisitable {
		visitable.VisitClass(f)
	}
}
