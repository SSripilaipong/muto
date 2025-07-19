package global

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestStringOps_stringToRunes(t *testing.T) {
	t.Run("should convert string to runes", func(t *testing.T) {
		x := base.NewNamedOneLayerObject(stringToRunesMutator.Name(), base.NewString("a\nμ"))
		y := stringToRunesMutator.Mutate(x)
		assert.Equal(t, base.NewConventionalList(base.NewRune('a'), base.NewRune('\n'), base.NewRune('μ')), y.Value())
	})
}

func TestStringOps_string(t *testing.T) {
	t.Run("should convert number to string", func(t *testing.T) {
		x := base.NewNamedOneLayerObject(stringMutator.Name(), base.NewNumberFromString("123"))
		y := stringMutator.Mutate(x)
		assert.Equal(t, base.NewString("123"), y.Value())
	})
}
