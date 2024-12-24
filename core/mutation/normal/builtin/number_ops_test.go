package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestNumberOps(t *testing.T) {
	t.Run("should add numbers", func(t *testing.T) {
		x := base.NewNamedOneLayerObject("+", []base.Node{base.NewNumberFromString("1"), base.NewNumberFromString("2")})
		y := addMutator.MutateByName("+", x)
		assert.Equal(t, base.NewNumberFromString("3"), y.Value())
	})

	t.Run("should subtract numbers", func(t *testing.T) {
		x := base.NewNamedOneLayerObject("-", []base.Node{base.NewNumberFromString("2"), base.NewNumberFromString("1")})
		y := subtractMutator.MutateByName("-", x)
		assert.Equal(t, base.NewNumberFromString("1"), y.Value())
	})
}
