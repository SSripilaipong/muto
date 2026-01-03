package pattern_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/syntaxtree"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	"github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestExtractHeadConjunctions_LayerCounting(t *testing.T) {
	t.Run("should return nil for non-object pattern", func(t *testing.T) {
		// pattern: f (bare class) - 0 layers
		p := syntaxtree.NewLocalClass("f")
		assert.Nil(t, pattern.ExtractHeadConjunctions(p))
	})

	t.Run("should return slice of length 1 for single layer", func(t *testing.T) {
		// pattern: f X - 1 layer
		p := pattern.NewDeterminantObject(
			syntaxtree.NewLocalClass("f"),
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
		)
		result := pattern.ExtractHeadConjunctions(p)
		assert.Equal(t, 1, len(result))
	})

	t.Run("should return slice of length 2 for two layers", func(t *testing.T) {
		// pattern: (f X) Y - 2 layers
		inner := pattern.NewDeterminantObject(
			syntaxtree.NewLocalClass("f"),
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
		)
		p := pattern.NewDeterminantObject(
			inner,
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("Y")}),
		)
		result := pattern.ExtractHeadConjunctions(p)
		assert.Equal(t, 2, len(result))
	})

	t.Run("should return slice of length 3 for three layers", func(t *testing.T) {
		// pattern: ((f X) Y) Z - 3 layers
		innermost := pattern.NewDeterminantObject(
			syntaxtree.NewLocalClass("f"),
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
		)
		inner := pattern.NewDeterminantObject(
			innermost,
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("Y")}),
		)
		p := pattern.NewDeterminantObject(
			inner,
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("Z")}),
		)
		result := pattern.ExtractHeadConjunctions(p)
		assert.Equal(t, 3, len(result))
	})

	t.Run("should not count DeterminantConjunction as extra layer", func(t *testing.T) {
		// pattern: (f X)^P Y - 2 layers (conjunction doesn't add)
		inner := pattern.NewDeterminantObject(
			syntaxtree.NewLocalClass("f"),
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
		)
		conj := pattern.NewDeterminantConjunction(inner, []base.Pattern{syntaxtree.NewVariable("P")})
		p := pattern.NewDeterminantObject(
			conj,
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("Y")}),
		)
		result := pattern.ExtractHeadConjunctions(p)
		assert.Equal(t, 2, len(result))
	})

	t.Run("should not count nested DeterminantConjunctions as extra layers", func(t *testing.T) {
		// pattern: ((f X)^P Y)^Q Z - 3 layers
		innermost := pattern.NewDeterminantObject(
			syntaxtree.NewLocalClass("f"),
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
		)
		innerConj := pattern.NewDeterminantConjunction(innermost, []base.Pattern{syntaxtree.NewVariable("P")})
		middle := pattern.NewDeterminantObject(
			innerConj,
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("Y")}),
		)
		outerConj := pattern.NewDeterminantConjunction(middle, []base.Pattern{syntaxtree.NewVariable("Q")})
		p := pattern.NewDeterminantObject(
			outerConj,
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("Z")}),
		)
		result := pattern.ExtractHeadConjunctions(p)
		assert.Equal(t, 3, len(result))
	})

	t.Run("should count 1 layer for bare class with conjunction", func(t *testing.T) {
		// pattern: f^P X - 1 layer
		conj := pattern.NewDeterminantConjunction(
			syntaxtree.NewLocalClass("f"),
			[]base.Pattern{syntaxtree.NewVariable("P")},
		)
		p := pattern.NewDeterminantObject(
			conj,
			pattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
		)
		result := pattern.ExtractHeadConjunctions(p)
		assert.Equal(t, 1, len(result))
	})
}
