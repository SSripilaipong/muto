package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestControl_doMutator(t *testing.T) {
	t.Run("should become last node", func(t *testing.T) {
		result := doMutator.Mutate(doMutatorName, base.NewNamedObject(doMutatorName, []base.Node{
			base.NewNumberFromString("123"), base.NewNumberFromString("456"),
		}))
		assert.Equal(t, base.NewNumberFromString("456"), result.Value())
	})

	t.Run("should mutate with one children", func(t *testing.T) {
		result := doMutator.Mutate(doMutatorName, base.NewNamedObject(doMutatorName, []base.Node{
			base.NewNumberFromString("123"),
		}))
		assert.Equal(t, base.NewNumberFromString("123"), result.Value())
	})

	t.Run("should mutate with more than 2 children", func(t *testing.T) {
		result := doMutator.Mutate(doMutatorName, base.NewNamedObject(doMutatorName, []base.Node{
			base.NewNumberFromString("123"), base.NewNumberFromString("456"), base.NewNumberFromString("789"),
		}))
		assert.Equal(t, base.NewNumberFromString("789"), result.Value())
	})

	t.Run("should not mutate when no children", func(t *testing.T) {
		result := doMutator.Mutate(doMutatorName, base.NewNamedObject(doMutatorName, []base.Node{}))
		assert.True(t, result.IsEmpty())
	})
}
