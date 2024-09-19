package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/optional"
)

func TestAnonymousObject_NormalMutation(t *testing.T) {
	t.Run("should bubble up to data node with children", func(t *testing.T) {
		node := NewAnonymousObject(NewString("hello"), []Node{NewNumberFromString("123"), NewString("world")})
		result := node.Mutate(nil)
		assert.Equal(t, NewDataObject([]Node{NewString("hello"), NewNumberFromString("123"), NewString("world")}), result.Value())
	})

	t.Run("should mutate head", func(t *testing.T) {
		node := NewAnonymousObject(NewNamedObject("hello", nil), nil)
		result := node.Mutate(newNormalMutationForTest(func(s string, object NamedObject) optional.Of[Node] {
			return optional.Value[Node](NewNamedObject("world", nil))
		}))
		assert.Equal(t, NewAnonymousObject(NewNamedObject("world", nil), nil), result.Value())
	})

	t.Run("should bubble up terminated head", func(t *testing.T) {
		node := NewAnonymousObject(NewNamedObject("hello", nil), nil)
		result := node.Mutate(newNormalMutationForTest(func(s string, object NamedObject) optional.Of[Node] { return optional.Empty[Node]() }))
		assert.Equal(t, NewNamedObject("hello", nil), result.Value())
	})

	t.Run("should bubble up terminated head with children", func(t *testing.T) {
		node := NewAnonymousObject(NewNamedObject("hello", []Node{NewString("a"), NewString("b")}), nil)
		result := node.Mutate(newNormalMutationForTest(func(s string, object NamedObject) optional.Of[Node] { return optional.Empty[Node]() }))
		assert.Equal(t, NewNamedObject("hello", []Node{NewString("a"), NewString("b")}), result.Value())
	})

	t.Run("should bubble up data to head to data node when children not exist", func(t *testing.T) {
		node := NewAnonymousObject(NewString("abc"), nil)
		result := node.Mutate(nil)
		assert.Equal(t, NewString("abc"), result.Value())
	})
}

type normalMutationForTest func(string, NamedObject) optional.Of[Node]

func newNormalMutationForTest(f func(string, NamedObject) optional.Of[Node]) normalMutationForTest {
	return f
}

func (f normalMutationForTest) Normal(name string, obj NamedObject) optional.Of[Node] {
	return f(name, obj)
}

func (f normalMutationForTest) Active(string, NamedObject) optional.Of[Node] {
	return optional.Empty[Node]()
}
