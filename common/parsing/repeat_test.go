package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionalGreedyRepeat(t *testing.T) {
	t.Run("should accept empty", func(t *testing.T) {
		p := OptionalGreedyRepeat(ToParser(ConsumeIf(func(s rune) bool { return s == 'a' })))
		r, k := p([]rune("")).Return()
		assert.True(t, r.IsOk())
		assert.Empty(t, k)
	})

	t.Run("should not consume other token", func(t *testing.T) {
		p := OptionalGreedyRepeat(ToParser(ConsumeIf(func(s rune) bool { return s == 'a' })))
		r, k := p([]rune("b")).Return()
		assert.True(t, r.IsOk())
		assert.Equal(t, k, []rune("b"))
	})
}
