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
}

func NewRightVariadic(name string, patternsFromLeft []NodeExtractor) RightVariadic {
	return RightVariadic{
		name:             name,
		nLeft:            len(patternsFromLeft),
		exactLeftPattern: NewExactNodeList(patternsFromLeft),
	}
}

func (p RightVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	nLeft := p.NLeft()
	if len(nodes) < nLeft {
		return optional.Empty[*parameter.Parameter]()
	}

	return optional.JoinFmap(parameter.WithVariadicVarMappings(parameter.NewVariadicVarMapping(p.Name(), nodes[nLeft:])))(
		p.ExactLeftPattern().Extract(nodes[:nLeft]),
	)
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

type IgnoredRightVariadic struct {
	nLeft            int
	exactLeftPattern ExactNodeList
}

func NewIgnoredRightVariadic(patternsFromLeft []NodeExtractor) IgnoredRightVariadic {
	return IgnoredRightVariadic{
		nLeft:            len(patternsFromLeft),
		exactLeftPattern: NewExactNodeList(patternsFromLeft),
	}
}

func (p IgnoredRightVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	nLeft := p.NLeft()
	if len(nodes) < nLeft {
		return optional.Empty[*parameter.Parameter]()
	}
	return p.ExactLeftPattern().Extract(nodes[:nLeft])
}

func (p IgnoredRightVariadic) NLeft() int {
	return p.nLeft
}

func (p IgnoredRightVariadic) ExactLeftPattern() ExactNodeList {
	return p.exactLeftPattern
}

var _ NodeListExtractor = IgnoredRightVariadic{}
