package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/optional"
)

func TestObject_WithObjectHead(t *testing.T) {
	t.Run("should mutate from mutation function", func(t *testing.T) {
		hello := NewRuleBasedClass("hello", newNormalMutationForTest(func(object Object) optional.Of[Node] {
			return optional.Value[Node](NewNamedOneLayerObject("world"))
		}))
		result := NewOneLayerObject(hello).Mutate()
		assert.Equal(t, NewNamedOneLayerObject("world"), result.Value())
	})

	t.Run("should bubble up data to head to data node when children not exist", func(t *testing.T) {
		node := NewOneLayerObject(NewString("abc"))
		result := node.Mutate()
		assert.True(t, result.IsEmpty())
	})
}

type normalMutationForTest func(Object) optional.Of[Node]

func newNormalMutationForTest(f func(Object) optional.Of[Node]) normalMutationForTest {
	return f
}

func (f normalMutationForTest) Normal(obj Object) optional.Of[Node] {
	return f(obj)
}

func (f normalMutationForTest) Active(Object) optional.Of[Node] {
	return optional.Empty[Node]()
}
