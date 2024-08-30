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
		assert.Equal(t, base.NewString("hello world"), program.Mutate(program.InitialObject()).Value())
	})

	t.Run("should resolve to number", func(t *testing.T) {
		program := BuildFromString(`main = 123.45`).Value()
		assert.Equal(t, base.NewNumber(datatype.NewNumber("123.45")), program.Mutate(program.InitialObject()).Value())
	})

	t.Run("should resolve to object", func(t *testing.T) {
		program := BuildFromString(`main = hello "world"`).Value()

		assert.Equal(t, base.NewObject(base.NewNamedClass("hello"), []base.Node{base.NewString("world")}), program.Mutate(program.InitialObject()).Value())
	})
}
