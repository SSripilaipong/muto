package extractor

import (
	"fmt"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type Conjunction struct {
	left  NodeExtractor
	right NodeExtractor
}

func NewConjunction(left, right NodeExtractor) NodeExtractor {
	return Conjunction{left: left, right: right}
}

func (c Conjunction) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	left := c.left.Extract(node)
	right := c.right.Extract(node)
	return optionalMergeParam(left, right)
}

func (c Conjunction) DisplayString() string {
	return fmt.Sprintf("%s^%s", DisplayString(c.left), DisplayString(c.right))
}

var _ NodeExtractor = Conjunction{}
