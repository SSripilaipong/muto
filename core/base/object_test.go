package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
)

func TestObject_WithObjectHead(t *testing.T) {
	t.Run("should mutate head", func(t *testing.T) {
		node := NewOneLayerObject(NewNamedOneLayerObject("hello", nil), nil)
		result := node.Mutate(newNormalMutationForTest(func(s string, object Object) optional.Of[Node] {
			return optional.Value[Node](NewNamedOneLayerObject("world", nil))
		}))
		assert.Equal(t, NewOneLayerObject(NewNamedOneLayerObject("world", nil), nil), result.Value())
	})

	t.Run("should bubble up terminated head", func(t *testing.T) {
		node := NewOneLayerObject(NewNamedOneLayerObject("hello", nil), nil)
		result := node.Mutate(newNormalMutationForTest(func(s string, object Object) optional.Of[Node] { return optional.Empty[Node]() }))
		assert.Equal(t, NewNamedOneLayerObject("hello", nil), result.Value())
	})

	t.Run("should bubble up terminated head with children", func(t *testing.T) {
		node := NewOneLayerObject(NewNamedOneLayerObject("hello", []Node{NewString("a"), NewString("b")}), nil)
		result := node.Mutate(newNormalMutationForTest(func(s string, object Object) optional.Of[Node] { return optional.Empty[Node]() }))
		assert.Equal(t, NewNamedOneLayerObject("hello", []Node{NewString("a"), NewString("b")}), result.Value())
	})

	t.Run("should bubble up data to head to data node when children not exist", func(t *testing.T) {
		node := NewOneLayerObject(NewString("abc"), nil)
		result := node.Mutate(nil)
		assert.Equal(t, NewString("abc"), result.Value())
	})
}

type normalMutationForTest func(string, Object) optional.Of[Node]

func newNormalMutationForTest(f func(string, Object) optional.Of[Node]) normalMutationForTest {
	return f
}

func (f normalMutationForTest) Normal(name string, obj Object) optional.Of[Node] {
	return f(name, obj)
}

func (f normalMutationForTest) Active(string, Object) optional.Of[Node] {
	return optional.Empty[Node]()
}
