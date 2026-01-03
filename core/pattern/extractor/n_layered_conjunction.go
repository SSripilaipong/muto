package extractor

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

// NLayeredConjunction wraps an inner extractor and applies level-based conjunctions.
// It extracts using the inner extractor first, then for each nesting level with conjunctions,
// it reconstructs the object at that level and applies the conjunction extractors.
type NLayeredConjunction struct {
	inner      NodeExtractor
	conjLayers [][]NodeExtractor
}

func NewNLayeredConjunction(inner NodeExtractor, conjLayers [][]NodeExtractor) NodeExtractor {
	if !hasAnyConjunctions(conjLayers) {
		return inner
	}
	return NLayeredConjunction{inner: inner, conjLayers: conjLayers}
}

func hasAnyConjunctions(conjLayers [][]NodeExtractor) bool {
	for _, layer := range conjLayers {
		if len(layer) > 0 {
			return true
		}
	}
	return false
}

func (c NLayeredConjunction) Extract(node base.Node) optional.Of[*parameter.Parameter] {
	// 1. Extract using inner extractor first
	result := c.inner.Extract(node)
	if result.IsEmpty() {
		return result
	}

	// 2. For each level, apply conjunction extractors
	for level, conjs := range c.conjLayers {
		if len(conjs) == 0 {
			continue
		}
		reconstructed := ReconstructNodeToLevel(node, level)
		for _, conjExtractor := range conjs {
			conjResult := conjExtractor.Extract(reconstructed)
			result = optionalMergeParam(result, conjResult)
			if result.IsEmpty() {
				return result
			}
		}
	}
	return result
}

// ReconstructNodeToLevel reconstructs a node at a specific nesting level.
// If the node is not an object, it returns the node unchanged.
func ReconstructNodeToLevel(node base.Node, level int) base.Node {
	if !base.IsObjectNode(node) {
		return node // Class/tag stays as-is
	}
	return ReconstructObjectToLevel(base.UnsafeNodeToObject(node), level)
}

// ReconstructObjectToLevel reconstructs the object at a specific nesting level.
// For an object with head h and params [[1],[2],[3]]:
//   - Level 0: h (just head)
//   - Level 1: (h 1) = NewCompoundObject(h, [[1]])
//   - Level 2: ((h 1) 2) = NewCompoundObject(h, [[1],[2]])
func ReconstructObjectToLevel(obj base.Object, level int) base.Node {
	if level == 0 {
		return obj.Head() // Just the head
	}
	params := obj.ParamChain().SliceUntilOrEmpty(level)
	if params.Size() == 0 {
		return obj.Head()
	}
	return base.NewCompoundObject(obj.Head(), params)
}

var _ NodeExtractor = NLayeredConjunction{}
