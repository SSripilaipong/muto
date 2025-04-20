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
}

func NewLeftVariadic(name string, rightPatterns []NodeExtractor) LeftVariadic {
	return LeftVariadic{
		name:              name,
		nRight:            len(rightPatterns),
		exactRightPattern: NewExactNodeList(rightPatterns),
	}
}

func (p LeftVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	nVariadic := len(nodes) - p.NRight()
	if nVariadic < 0 {
		return optional.Empty[*parameter.Parameter]()
	}

	return optional.JoinFmap(parameter.WithVariadicVarMappings(parameter.NewVariadicVarMapping(p.Name(), nodes[:nVariadic])))(
		p.ExactRightPattern().Extract(nodes[nVariadic:]),
	)
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

type IgnoredLeftVariadic struct {
	nRight            int
	exactRightPattern ExactNodeList
}

func NewIgnoredLeftVariadic(rightPatterns []NodeExtractor) IgnoredLeftVariadic {
	return IgnoredLeftVariadic{
		nRight:            len(rightPatterns),
		exactRightPattern: NewExactNodeList(rightPatterns),
	}
}

func (p IgnoredLeftVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	nVariadic := len(nodes) - p.NRight()
	if nVariadic < 0 {
		return optional.Empty[*parameter.Parameter]()
	}

	return p.ExactRightPattern().Extract(nodes[nVariadic:])
}

func (p IgnoredLeftVariadic) NRight() int {
	return p.nRight
}

func (p IgnoredLeftVariadic) ExactRightPattern() ExactNodeList {
	return p.exactRightPattern
}

var _ NodeListExtractor = IgnoredLeftVariadic{}
