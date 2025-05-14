package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

func TestObjectOps_try(t *testing.T) {
	try := newTryMutator()

	t.Run("should call global mutator with the concat object", func(t *testing.T) {
		var calledObject base.Object
		f := base.NewClass("f", newNormalMutationFunc(func(obj base.Object) optional.Of[base.Node] {
			calledObject = obj
			return optional.Empty[base.Node]()
		}))
		try.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{f, base.NewString("abc"), base.NewString("def")}))
		assert.True(t, base.NodeEqual(
			base.NewOneLayerObject(base.NewUnlinkedClass("f"), []base.Node{base.NewString("abc"), base.NewString("def")}),
			calledObject,
		))
	})

	t.Run("should become value of the mutation result", func(t *testing.T) {
		f := base.NewClass("f", newNormalMutationFunc(func(obj base.Object) optional.Of[base.Node] {
			return optional.Value[base.Node](base.NewNumberFromString("123"))
		}))
		result := try.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{f, base.NewString("abc")}))
		assert.True(t, base.NodeEqual(
			base.NewOneLayerObject(base.ValueTag, []base.Node{base.NewNumberFromString("123")}), result.Value()),
		)
	})

	t.Run("should become empty if mutation does not occur", func(t *testing.T) {
		result := try.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{base.NewUnlinkedClass("f"), base.NewString("abc")}))
		assert.Equal(t, base.EmptyTag, result.Value())
	})

	t.Run("should not mutate if the number of children less than 2", func(t *testing.T) {
		result := try.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{base.NewUnlinkedClass("f")}))
		assert.True(t, result.IsEmpty())
	})
}
