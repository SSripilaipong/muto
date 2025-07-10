package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type RightVariadic struct {
	name             string
	nLeft            int
	exactLeftPattern ExactNodeList
	variadicVar      NodeListExtractor
}

func NewRightVariadic(name string, variadicVar NodeListExtractor, patternsFromLeft []NodeExtractor) RightVariadic {
	return RightVariadic{
		name:             name,
		nLeft:            len(patternsFromLeft),
		exactLeftPattern: NewExactNodeList(patternsFromLeft),
		variadicVar:      variadicVar,
	}
}

func (p RightVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	nLeft := p.NLeft()
	if len(nodes) < nLeft {
		return optional.Empty[*parameter.Parameter]()
	}

	variadicParams := p.variadicVar.Extract(nodes[nLeft:])
	fixedParams := p.ExactLeftPattern().Extract(nodes[:nLeft])
	return optional.MergeFn(parameter.Merge)(variadicParams, fixedParams)
}

func (p RightVariadic) Name() string {
	return p.name
}

func (p RightVariadic) NLeft() int {
	return p.nLeft
}

func (p RightVariadic) ExactLeftPattern() ExactNodeList {
	return p.exactLeftPattern
}

var _ NodeListExtractor = RightVariadic{}
