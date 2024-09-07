package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"muto/common/optional"
)

func TestAnonymousObject_MutateWithObjMutateFunc(t *testing.T) {
	t.Run("should bubble up to data node", func(t *testing.T) {
		node := NewAnonymousObject(NewString("hello"), nil)
		result := node.MutateWithObjMutateFunc(nil)
		assert.Equal(t, NewDataObject([]Node{NewString("hello")}), result.Value())
	})

	t.Run("should bubble up to data node with children", func(t *testing.T) {
		node := NewAnonymousObject(NewString("hello"), []Node{NewNumberFromString("123"), NewString("world")})
		result := node.MutateWithObjMutateFunc(nil)
		assert.Equal(t, NewDataObject([]Node{NewString("hello"), NewNumberFromString("123"), NewString("world")}), result.Value())
	})

	t.Run("should mutate head", func(t *testing.T) {
		node := NewAnonymousObject(NewNamedObject("hello", nil), nil)
		result := node.MutateWithObjMutateFunc(func(s string, object NamedObject) optional.Of[Node] {
			return optional.Value[Node](NewNamedObject("world", nil))
		})
		assert.Equal(t, NewAnonymousObject(NewNamedObject("world", nil), nil), result.Value())
	})

	t.Run("should bubble up terminated head", func(t *testing.T) {
		node := NewAnonymousObject(NewNamedObject("hello", nil), nil)
		result := node.MutateWithObjMutateFunc(func(s string, object NamedObject) optional.Of[Node] { return optional.Empty[Node]() })
		assert.Equal(t, NewNamedObject("hello", nil), result.Value())
	})

	t.Run("should bubble up terminated head with children", func(t *testing.T) {
		node := NewAnonymousObject(NewNamedObject("hello", []Node{NewString("a"), NewString("b")}), nil)
		result := node.MutateWithObjMutateFunc(func(s string, object NamedObject) optional.Of[Node] { return optional.Empty[Node]() })
		assert.Equal(t, NewNamedObject("hello", []Node{NewString("a"), NewString("b")}), result.Value())
	})
}