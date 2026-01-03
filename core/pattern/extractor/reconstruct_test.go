package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestReconstructObjectToLevel(t *testing.T) {
	t.Run("should return head at level 0", func(t *testing.T) {
		// Object: (f 1 2)
		obj := base.NewNamedOneLayerObject("f", base.NewNumberFromString("1"), base.NewNumberFromString("2"))

		result := ReconstructObjectToLevel(obj, 0)

		assert.True(t, base.NewUnlinkedRuleBasedClass("f").Equals(base.UnsafeNodeToRuleBasedClass(result)))
	})

	t.Run("should return object with first param layer at level 1", func(t *testing.T) {
		// Object: (f 1) 2 = CompoundObject(f, [[1], [2]])
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
			}),
		)

		result := ReconstructObjectToLevel(obj, 1)

		expected := base.NewNamedOneLayerObject("f", base.NewNumberFromString("1"))
		assert.True(t, expected.Equals(base.UnsafeNodeToObject(result)))
	})

	t.Run("should return object with first two param layers at level 2", func(t *testing.T) {
		// Object: ((f 1) 2) 3 = CompoundObject(f, [[1], [2], [3]])
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
				{base.NewNumberFromString("3")},
			}),
		)

		result := ReconstructObjectToLevel(obj, 2)

		expected := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
			}),
		)
		assert.True(t, expected.Equals(base.UnsafeNodeToObject(result)))
	})

	t.Run("should handle object with multiple children per level", func(t *testing.T) {
		// Object: (f 1 2) 3 4 = CompoundObject(f, [[1, 2], [3, 4]])
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1"), base.NewNumberFromString("2")},
				{base.NewNumberFromString("3"), base.NewNumberFromString("4")},
			}),
		)

		result := ReconstructObjectToLevel(obj, 1)

		expected := base.NewNamedOneLayerObject("f", base.NewNumberFromString("1"), base.NewNumberFromString("2"))
		assert.True(t, expected.Equals(base.UnsafeNodeToObject(result)))
	})
}

func TestReconstructNodeToLevel(t *testing.T) {
	t.Run("should return class unchanged", func(t *testing.T) {
		class := base.NewUnlinkedRuleBasedClass("f")

		result := ReconstructNodeToLevel(class, 1)

		assert.True(t, class.Equals(base.UnsafeNodeToRuleBasedClass(result)))
	})

	t.Run("should return tag unchanged", func(t *testing.T) {
		tag := base.NewTag("abc")

		result := ReconstructNodeToLevel(tag, 1)

		assert.Equal(t, tag, result)
	})

	t.Run("should delegate to ReconstructObjectToLevel for objects", func(t *testing.T) {
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
			}),
		)

		result := ReconstructNodeToLevel(obj, 1)

		expected := base.NewNamedOneLayerObject("f", base.NewNumberFromString("1"))
		assert.True(t, expected.Equals(base.UnsafeNodeToObject(result)))
	})
}
