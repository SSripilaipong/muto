package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type LeftVariadic struct {
	name              string
	nRight            int
	exactRightPattern ExactNodeList
	variadicVar       NodeListExtractor
}

func NewLeftVariadic(name string, variadicVar NodeListExtractor, rightPatterns []NodeExtractor) LeftVariadic {
	return LeftVariadic{
		name:              name,
		nRight:            len(rightPatterns),
		exactRightPattern: NewExactNodeList(rightPatterns),
		variadicVar:       variadicVar,
	}
}

func (p LeftVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	nVariadic := len(nodes) - p.NRight()
	if nVariadic < 0 {
		return optional.Empty[*parameter.Parameter]()
	}

	variadicParams := p.variadicVar.Extract(nodes[:nVariadic])
	fixedParams := p.ExactRightPattern().Extract(nodes[nVariadic:])
	return optional.MergeFn(parameter.Merge)(variadicParams, fixedParams)
}

func (p LeftVariadic) Name() string {
	return p.name
}

func (p LeftVariadic) NRight() int {
	return p.nRight
}

func (p LeftVariadic) ExactRightPattern() ExactNodeList {
	return p.exactRightPattern
}

var _ NodeListExtractor = LeftVariadic{}
