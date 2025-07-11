package global

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestStringOps(t *testing.T) {
	t.Run("should convert number to string", func(t *testing.T) {
		x := base.NewNamedOneLayerObject("string", []base.Node{base.NewNumberFromString("123")})
		y := stringMutator.MutateByName("string", x)
		assert.Equal(t, base.NewString("123"), y.Value())
	})
}
