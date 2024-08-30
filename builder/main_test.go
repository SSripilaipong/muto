package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"phi-lang/core/base"
	"phi-lang/core/base/datatype"
)

func TestBuildFromString(t *testing.T) {
	t.Run("should resolve to string", func(t *testing.T) {
		program := BuildFromString(`main = "hello world"`).Value()
		assert.Equal(t, base.NewString("hello world"), mutateN(1, program))
	})

	t.Run("should resolve to number", func(t *testing.T) {
		program := BuildFromString(`main = 123.45`).Value()
		assert.Equal(t, base.NewNumber(datatype.NewNumber("123.45")), mutateN(1, program))
	})

	t.Run("should resolve to object", func(t *testing.T) {
		program := BuildFromString(`main = hello "world"`).Value()
		assert.Equal(t, base.NewObject(base.NewNamedClass("hello"), []base.Node{base.NewString("world")}), mutateN(1, program))
	})

	t.Run("should resolve variable", func(t *testing.T) {
		program := BuildFromString(`
hello X = X
main = hello "world"
`).Value()
		assert.Equal(t, base.NewString("world"), mutateN(2, program))
	})

	t.Run("should resolve builtin add object", func(t *testing.T) {
		program := BuildFromString(`main = + 1 2`).Value()
		assert.Equal(t, base.NewNumber(datatype.NewNumber("3")), mutateN(3, program))
	})

	t.Run("should resolve builtin concat object", func(t *testing.T) {
		program := BuildFromString(`main = ++ "hello" " world"`).Value()
		assert.Equal(t, base.NewString("hello world"), mutateN(3, program))
	})
}

func mutateN(n int, program Program) base.Node {
	if n == 0 {
		return program.InitialObject()
	}
	return program.Mutate(mutateN(n-1, program))
}
