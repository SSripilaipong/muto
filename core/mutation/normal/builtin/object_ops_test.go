package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

func TestObjectOps_try(t *testing.T) {
	t.Run("should call global mutator with the concat object", func(t *testing.T) {
		m := newTryMutator()
		var calledObject base.Object
		m.SetGlobalMutator(mutator.NameBoundedFunc(func(name string, obj base.Object) optional.Of[base.Node] {
			calledObject = obj
			return optional.Empty[base.Node]()
		}))
		m.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{base.NewClass("f"), base.NewString("abc"), base.NewString("def")}))
		assert.Equal(t, base.NewOneLayerObject(base.NewClass("f"), []base.Node{base.NewString("abc"), base.NewString("def")}), calledObject)
	})

	t.Run("should become value of the mutation result", func(t *testing.T) {
		m := newTryMutator()
		m.SetGlobalMutator(mutator.NameBoundedFunc(func(name string, obj base.Object) optional.Of[base.Node] {
			return optional.Value[base.Node](base.NewNumberFromString("123"))
		}))
		result := m.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{base.NewClass("f"), base.NewString("abc")}))
		assert.Equal(t, base.NewOneLayerObject(base.ValueTag, []base.Node{base.NewNumberFromString("123")}), result.Value())
	})

	t.Run("should become empty if mutation does not occur", func(t *testing.T) {
		m := newTryMutator()
		m.SetGlobalMutator(mutator.NameBoundedFunc(func(name string, obj base.Object) optional.Of[base.Node] {
			return optional.Empty[base.Node]()
		}))
		result := m.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{base.NewClass("f"), base.NewString("abc")}))
		assert.Equal(t, base.NewOneLayerObject(base.EmptyTag, []base.Node{}), result.Value())
	})

	t.Run("should not mutate if the number of children less than 2", func(t *testing.T) {
		m := newTryMutator()
		m.SetGlobalMutator(mutator.NameBoundedFunc(func(name string, obj base.Object) optional.Of[base.Node] {
			return optional.Value[base.Node](base.NewNumberFromString("123"))
		}))
		result := m.MutateByName(tryMutatorName, base.NewNamedOneLayerObject(tryMutatorName, []base.Node{base.NewClass("f")}))
		assert.True(t, result.IsEmpty())
	})
}
