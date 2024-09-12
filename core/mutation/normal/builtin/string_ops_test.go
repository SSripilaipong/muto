package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"muto/core/base"
)

func TestStringOps(t *testing.T) {
	t.Run("should convert number to string", func(t *testing.T) {
		x := base.NewNamedObject("string", []base.Node{base.NewNumberFromString("123")})
		y := stringMutator.Mutate(x)
		assert.Equal(t, base.NewString("123"), y.Value())
	})
}