package pattern

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type Object interface {
	base.Pattern
	Head() base.Pattern
	ParamPart() ParamPart
}

func UnsafePatternToObject(p base.Pattern) Object {
	return p.(Object)
}
