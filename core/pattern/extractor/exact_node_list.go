package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type ExactNodeList struct {
	patterns []NodeExtractor
}

func NewExactNodeList(patternsFromLeft []NodeExtractor) ExactNodeList {
	return ExactNodeList{patterns: patternsFromLeft}
}

func (p ExactNodeList) Extract(nodes []base.Node) optional.Of[*parameter.Parameter] {
	if len(nodes) != p.NConsumed() {
		return optional.Empty[*parameter.Parameter]()
	}

	patternsFromLeft := p.PatternsFromLeft()
	mutation := parameter.New()
	for i, node := range nodes {
		e := patternsFromLeft[i].Extract(node)
		if e.IsEmpty() {
			return optional.Empty[*parameter.Parameter]()
		}
		m := mutation.Merge(e.Value())
		if m.IsEmpty() {
			return optional.Empty[*parameter.Parameter]()
		}
		mutation = m.Value()
	}

	return optional.Value(mutation)
}

func (p ExactNodeList) PatternsFromLeft() []NodeExtractor {
	return p.patterns
}

func (p ExactNodeList) NConsumed() int {
	return len(p.PatternsFromLeft())
}

var _ NodeListExtractor = ExactNodeList{}
