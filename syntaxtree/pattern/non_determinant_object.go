package pattern

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

type NonDeterminantObject struct {
	head          base.Pattern
	params        ParamPart
	nonNestedHead base.Pattern
}

func (NonDeterminantObject) PatternType() base.PatternType { return base.PatternTypeObject }

func (p NonDeterminantObject) Head() base.Pattern {
	return p.head
}

func (p NonDeterminantObject) ParamPart() ParamPart {
	return p.params
}

func (p NonDeterminantObject) ParamParts() []ParamPart {
	return slc.Pure(p.params)
}

func NewNonDeterminantObject(head base.Pattern, params ParamPart) NonDeterminantObject {
	return NonDeterminantObject{
		head:          head,
		params:        params,
		nonNestedHead: ExtractNonObjectHead(head),
	}
}
