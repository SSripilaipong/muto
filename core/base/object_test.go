package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
)

func TestObject_WithObjectHead(t *testing.T) {
	t.Run("should mutate from mutation function", func(t *testing.T) {
		node := NewCompoundObject(NewClass("hello"), NewParamChain([][]Node{{}}))
		result := node.Mutate(newNormalMutationForTest(func(s string, object Object) optional.Of[Node] {
			return optional.Value[Node](NewNamedOneLayerObject("world", nil))
		}))
		assert.Equal(t, NewCompoundObject(NewClass("world"), NewParamChain([][]Node{nil})), result.Value())
	})

	t.Run("should bubble up data to head to data node when children not exist", func(t *testing.T) {
		node := NewOneLayerObject(NewString("abc"), nil)
		result := node.Mutate(nil)
		assert.True(t, result.IsEmpty())
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
