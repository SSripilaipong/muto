package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"muto/core/base"
)

func TestBuildFromString(t *testing.T) {
	t.Run("should resolve to string", func(t *testing.T) {
		program := BuildFromString(`main = "hello world"`).Value()
		assert.Equal(t, base.NewString("hello world"), execute(program))
	})

	t.Run("should resolve to number", func(t *testing.T) {
		program := BuildFromString(`main = 123.45`).Value()
		assert.Equal(t, base.NewNumberFromString("123.45"), execute(program))
	})

	t.Run("should resolve to object", func(t *testing.T) {
		program := BuildFromString(`main = hello "world"`).Value()
		assert.Equal(t, base.NewNamedObject("hello", []base.Node{base.NewString("world")}), execute(program))
	})

	t.Run("should resolve variable", func(t *testing.T) {
		program := BuildFromString(`
hello X = X
main = hello "world"
`).Value()
		assert.Equal(t, base.NewString("world"), execute(program))
	})

	t.Run("should resolve builtin add object", func(t *testing.T) {
		program := BuildFromString(`main = + 1 2`).Value()
		assert.Equal(t, base.NewNumberFromString("3"), execute(program))
	})

	t.Run("should resolve builtin concat object", func(t *testing.T) {
		program := BuildFromString(`main = ++ "hello" " world"`).Value()
		assert.Equal(t, base.NewString("hello world"), execute(program))
	})

	t.Run("should match rule by string value", func(t *testing.T) {
		program := BuildFromString(`
hello "a" = 1
hello "b" = 2
main = hello "b"
`).Value()
		assert.Equal(t, base.NewNumberFromString("2"), execute(program))
	})

	t.Run("should match rule by number value", func(t *testing.T) {
		program := BuildFromString(`
hello 1 = "hello"
hello 2 = "world"
hello 3 = "muto"
main = hello 2
`).Value()
		assert.Equal(t, base.NewString("world"), execute(program))
	})

	t.Run("should resolve nested object", func(t *testing.T) {
		program := BuildFromString(`
hello (f X Y) = X
main = hello (f "abc" 123)
`).Value()
		assert.Equal(t, base.NewString("abc"), execute(program))
	})

	t.Run("should resolve with post-order mutation", func(t *testing.T) {
		program := BuildFromString(`main = ++ "hello " (string (+ 3 1)) (string (+ 1 1))`).Value()
		assert.Equal(t, base.NewString("hello 42"), execute(program))
	})

	t.Run("should resolve variadic variable", func(t *testing.T) {
		program := BuildFromString(`f X Xs... = g Xs...
main = f 1 2 3
`).Value()
		assert.Equal(t, base.NewNamedObject("g", []base.Node{base.NewNumberFromString("2"), base.NewNumberFromString("3")}), execute(program))
	})

	t.Run("should match children strictly for nested pattern", func(t *testing.T) {
		program := BuildFromString(`g (f 1) = 555
main = g (f 1 2)
`).Value()
		assert.Equal(t, base.NewNamedObject("g", []base.Node{base.NewNamedObject("f", []base.Node{base.NewNumberFromString("1"), base.NewNumberFromString("2")}).ConfirmTermination()}), execute(program))
	})

	t.Run("should resolve to data object when there are children left", func(t *testing.T) {
		program := BuildFromString(`f X = 999
main = f 1 2
`).Value()
		assert.Equal(t, base.NewDataObject([]base.Node{base.NewNumberFromString("999"), base.NewNumberFromString("2")}), execute(program))
	})
}

func execute(program Program) base.Node {
	return program.MutateUntilTerminated(program.InitialObject())
}
