package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type ImplicitRightVariadic struct {
	nLeft            int
	exactLeftPattern ExactNodeList
}

func NewImplicitRightVariadic(patternsFromLeft []NodeExtractor) ImplicitRightVariadic {
	return ImplicitRightVariadic{
		nLeft:            len(patternsFromLeft),
		exactLeftPattern: NewExactNodeList(patternsFromLeft),
	}
}

func (p ImplicitRightVariadic) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	nLeft := p.NLeft()
	if len(nodes) < nLeft {
		return optional.Empty[*parameter.Parameter]()
	}

	param := optional.Value(parameter.New())
	if len(nodes) > 0 {
		param = p.ExactLeftPattern().Extract(nodes[:nLeft])
	}
	return optional.Fmap(parameter.WithRemainingChildren(nodes[nLeft:]))(param)
}

func (p ImplicitRightVariadic) ExactLeftPattern() ExactNodeList {
	return p.exactLeftPattern
}

func (p ImplicitRightVariadic) NLeft() int {
	return p.nLeft
}

var _ NodeListExtractor = ImplicitRightVariadic{}
