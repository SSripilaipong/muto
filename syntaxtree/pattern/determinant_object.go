package pattern

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type DeterminantObject struct {
	head          base.Pattern
	params        ParamPart
	nonNestedHead base.Class
}

func (DeterminantObject) DeterminantType() base.DeterminantType { return base.DeterminantTypeObject }

func (DeterminantObject) PatternType() base.PatternType { return base.PatternTypeObject }

func (p DeterminantObject) Head() base.Pattern {
	return p.head
}

func (p DeterminantObject) ParamPart() ParamPart {
	return p.params
}

func (p DeterminantObject) ObjectName() string {
	return p.nonNestedHead.Name()
}

func NewDeterminantObject(head base.Determinant, params ParamPart) DeterminantObject {
	nonNestedHead := ExtractNonObjectHead(head)
	if !base.IsPatternTypeClass(nonNestedHead) {
		panic("assertion failed: determinant head should be class only")
	}
	return DeterminantObject{
		head:          head,
		params:        params,
		nonNestedHead: base.UnsafePatternToClass(nonNestedHead),
	}
}
